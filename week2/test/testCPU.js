const rp = require('request-promise');

const host = 3000

async function request () {
    console.log("let't go")
    await rp(`http://localhost:${host}/cpu`)
    console.log('done')
}

(async () => {
    await Promise.all([
        request(),
        request(),
        request(),
        request(),
        request(),
        request(),
    ])
})()