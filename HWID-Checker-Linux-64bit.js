const si = require('systeminformation');

si.uuid().then(id => hwid = id.os.toString()).then(x => {
    console.log("HWID-Checker")
    console.log("Your HWID: "+x)


});
setInterval(() => {
    let x = "x"
}, 15000);