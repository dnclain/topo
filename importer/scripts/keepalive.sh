#! /bin/bash

# update scripts files in /tmp
/bin/bash ./update.sh

python3 -c $'import time\nwhile True:\n     time.sleep(3600)'

return 0