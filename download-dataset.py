# 🚨 PY 3.9 scripts
# stlib
import os
import re
import sys
import shutil
import argparse
from typing import Callable
from progressbar.bar import ProgressBar
import requests
from progressbar import progressbar
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
            print("Sorry, should be one one [%s]" % '/'.join(expected))
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
    regex = re.compile(rf'.*{_filter}\.7z$')

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

        print('> 📡 Downloading %s -> %s' % (url, out_dir))

        # download
        r = requests.get(url, stream=True)
        # print('After %s', r.elapsed)
        if r.status_code == requests.codes.ok:
            _file_size = int(r.headers.get('Content-Length'))
            _chunk_size = 2**16
            print("- File size : %s Mb" % (_file_size/1024))
            with open(out_file, 'wb') as f:
                _i = 0
                _bar = ProgressBar(max_value=_file_size)
                for data in r.iter_content(chunk_size=_chunk_size):
                    _i += len(data)
                    f.write(data)
                    _bar.update(_i)
        else:
            return '<' + filename + " téléchargement en échec"
        return '<' + filename + " téléchargement réussi"
    return download_url


if __name__ == '__main__':
    # Args definition
    _parser = argparse.ArgumentParser(
        description='BDTOPOV3 Database Downloader')
    _parser.add_argument('filter', type=str, nargs='?', default='SHP',
                         help="A regular expression to filter out some links (SQL, SHP, GPK for example) (default: SHP)")
    _parser.add_argument('--output', type=str, nargs='?', default='topo-express',
                         help="The output directory. Will be created if it does not exist yet (default: topo-express)")

    _args = _parser.parse_args()
    if _args.filter != '':
        print('Filtering %s' % _args.filter)
    _found_links = filter_links(URL, only=_args.filter)

    for i in _found_links:
        print(i)
    print('------')
    print(f'Found {len(_found_links)} links')

    _response_continue: str = expects('Should I continue ?', ['y', 'n'])
    if _response_continue == 'n':
        print("❌ abort")
        sys.exit()

    _out_dir: str = check_output(_args.output)

    print("✅ OK, let's Download...")

    # results = ThreadPool(5).imap_unordered()
    _downloader = gen_url_downloader(_out_dir)
    for _i in _found_links:
        print(_downloader(_i))
