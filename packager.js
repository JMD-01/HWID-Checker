const { exec } = require('pkg')

async function test(){
    await exec([ 'HWID-Checker-Linux-64bit.js', '--target', 'node12-linux-x64', '--output', 'HWID-Checker-Linux-64bit' ])
}
test();