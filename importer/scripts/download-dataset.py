# 🚨 PY 3.9 scripts
# stlib
import os
import re
import sys
import shutil
import argparse
from typing import Callable
import requests
from tqdm import tqdm
from multiprocessing.pool import ThreadPool
# external
from bs4 import BeautifulSoup

# Mirror
URL = 'https://files.opendatarchives.fr/professionnels.ign.fr/bdtopo/latest/'


def expects(prompt: str, expected: list[str]) -> str:
    '''
    Prompt and filter response.
    '''
    if not isinstance(expected, (list, tuple)):
        expected = [expected]

    _i = None
    while _i is None or _i not in expected:
        if _i is not None:
            print("Sorry, should be one of [%s]" % '/'.join(expected))
        _i = input("%s [%s] : " % (prompt, '/'.join(expected)))

    print('Your choice :', _i)

    return _i


def check_output(path: str) -> None:
    '''
    Ensure output dir exists.
    May remove the existing one.
    '''
    print("📁 Output dir is :", _args.output)

    if os.path.exists(path):
        print(
            "🚨 It already exists. You may have file conflicts.")
        _response_continue: str = expects(
            'Remove ?', ['y', 'n', 'q'])
        if _response_continue == 'n':
            print("✅ Keep current folder.")
            return path
        if _response_continue == 'q':
            print("❌ abort")
            sys.exit()
        print("🥊 Delete...", end=None)
        shutil.rmtree(path)
        print("Done.")

    print("💾 Fresh output folder creation...")
    os.mkdir(path)
    print("Done.")

    return path


def extract_all_links(url_site) -> list[str]:
    html = requests.get(url_site).text
    soup = BeautifulSoup(html, 'html.parser').find_all('a')
    links = [link.get('href') for link in soup]
    return links


def filter_links(url: str, only=''):
    _filter = only + (r'.*' if only != '' else '')
    regex = re.compile(rf'^.*{_filter}\.7z.*$')

    all_links = extract_all_links(url)
    all_links = [i for i in all_links if regex.match(i)]
    all_links = [url+x for x in all_links]

    return all_links


def gen_url_downloader(out_dir: str) -> Callable[[str], str]:
    def download_url(url: str) -> str:
        # extract the file name from URL.
        file_name_start_pos = url.rfind("/") + 1
        filename = url[file_name_start_pos:]
        out_file = os.path.join(out_dir, filename)

        print('> ⬇ 🗺 %s' % url)

        # download
        r = requests.get(url, stream=True)
        # print('After %s', r.elapsed)
        if r.status_code == requests.codes.ok:
            _file_size = int(r.headers.get('Content-Length'))
            _chunk_size = 2**16
            tqdm.write("- 📥 🗺 %s : %s bytes" % (filename, _file_size))
            with open(out_file, 'wb') as f:
                _i = 0
                with tqdm(total=_file_size, unit='B', unit_scale=True, unit_divisor=1024) as _bar:
                    for data in r.iter_content(chunk_size=_chunk_size):
                        f.write(data)
                        _bar.update(len(data))
        else:
            return "< ❌ %s failed" % filename
        return "< 🎉 %s success (%s bytes) " % (filename, _i)
    return download_url


if __name__ == '__main__':
    # Args definition
    _parser = argparse.ArgumentParser(
        description='BDTOPOV3 Database Downloader')
    _parser.add_argument('filter', type=str, nargs='?', default='SHP',
                         help="A regular expression to filter out some links (SQL, SHP, GPK for example) (default: SHP)")
    _parser.add_argument('--output', type=str, nargs='?', default='topo-express',
                         help="The output directory. Will be created if it does not exist yet (default: topo-express)")
    _parser.add_argument('--threads', type=int, nargs='?', default=1,
                         help="Number of simulteanous downloads (default: 1)")

    _args = _parser.parse_args()

    _idx_url = URL
    _out_dir: str = check_output(_args.output)
    _threads: int = _args.threads

    # Utilisation d'une autre URL de téléchargement que l'URL par défaut
    if "DOWNLOAD_URL" in os.environ:
        if os.environ['DOWNLOAD_URL'] != "":
            _idx_url = os.environ['DOWNLOAD_URL']

    # Utilisation d'un autre maximum de téléchargements en parallèle que celui par défaut
    if "MAX_PARALLEL_DL" in os.environ:
        if int(os.environ['MAX_PARALLEL_DL']) != 0:
            _threads = int(os.environ['MAX_PARALLEL_DL'])
        else:
            print(
                "🚧 new value of MAX_PARALLEL_DL is not in integer. Ignored.", file=sys.stderr)

    if _args.filter != '':
        print('Filtering %s' % _args.filter)
    _found_links = filter_links(_idx_url, only=_args.filter)

    # Pour un simple test on se limite à une seule archive
    if "TEST_IMPORTER" in os.environ:
        if int(os.environ['TEST_IMPORTER']) != 0:
            print(
                "🚨 Testing MODE. Only the first link will be downloaded.", file=sys.stderr)
            _found_links = _found_links[:1]
        else:
            print(
                "🚧 new value of TEST_IMPORTER is not in integer. Ignored.", file=sys.stderr)

    for i in _found_links:
        print(i)
    print('------')
    print(f'Found {len(_found_links)} links')

    _response_continue: str = expects('Should I continue ?', ['y', 'n'])
    if _response_continue == 'n':
        print("❌ abort")
        sys.exit()

    print("✅ OK, let's Download...")

    # 1. regular
    #_downloader = gen_url_downloader(_out_dir)
    # for _i in _found_links:
    #    print(_downloader(_i))
    # 2. Parallel
    _pool = ThreadPool(_threads)
    results = _pool.imap_unordered(
        gen_url_downloader(_out_dir), _found_links)

    for r in results:
        # sys.stdout.flush()
        tqdm.write(r)
    _pool.close()
    _pool.join()

    tqdm.write("🎉 All task done")
