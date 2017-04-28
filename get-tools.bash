NODE_JS_VER=6.10.2
TOOLS_DIR=~/tool
NODE_JS_DIR=${TOOLS_DIR}/nodejs
PLATFORM=x86 #x64 TODO automate
NODE=node-v${NODE_JS_VER}-linux-${PLATFORM}

mkdir ${NODE_JS_DIR}

wget https://nodejs.org/dist/v${NODE_JS_VER}/${NODE}.tar.xz -P ${NODE_JS_DIR}

tar xvfJ ${NODE_JS_DIR}/${NODE}.tar.xz -C ${NODE_JS_DIR}

SETUP_ENV_FILE=setup-env.bash

echo export PATH=\$PATH:${NODE_JS_DIR}/${NODE}/bin > ${SETUP_ENV_FILE}
echo 'echo Node version: `node -v`'  >> ${SETUP_ENV_FILE}
echo 'echo NPM version: `npm -v`'  >> ${SETUP_ENV_FILE}

export PATH=$PATH:${NODE_JS_DIR}/${NODE}/bin

npm list -g
du -sh ${NODE_JS_DIR}/${NODE}/lib
npm install -g @angular/cli
npm list -g | grep angular
du -sh ${NODE_JS_DIR}/${NODE}/lib

