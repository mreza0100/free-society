DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR/../../
ROOT=$(pwd)





export ROOT=$ROOT
export MODE=dev
export SALT="saltsalt___salt"
export SECRET_KEY="ap:OUWE#@#9iwjd@u3wj20i2erakwjdfAOJGF_@!I"
export GIN_MODE=debug
