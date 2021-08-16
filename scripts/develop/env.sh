DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR/../../
ROOT=$(pwd)


export FEED_CONFIG_FILE_PATH=${ROOT}/services/feed/configs/configs.json
export HELLGATE_CONFIG_FILE_PATH=${ROOT}/services/hellgate/configs/configs.json
export NOTIFICATION_CONFIG_FILE_PATH=${ROOT}/services/notification/configs/configs.json
export POST_CONFIG_FILE_PATH=${ROOT}/services/post/configs/configs.json
export RELATION_CONFIG_FILE_PATH=${ROOT}/services/relation/configs/configs.json
export SECURITY_CONFIG_FILE_PATH=${ROOT}/services/security/configs/configs.json
export USER_CONFIG_FILE_PATH=${ROOT}/services/user/configs/configs.json


export ROOT=$ROOT
export MODE=dev
export SALT="saltsalt___salt"
export SECRET_KEY="ap:OUWE#@#9iwjd@u3wj20i2erakwjdfAOJGF_@!I"
export GIN_MODE=debug
