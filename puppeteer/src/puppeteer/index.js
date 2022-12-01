const puppeteer = require('puppeteer');

async function getScreenshot(url, width, height) {
    const browser = await puppeteer.launch({args: ['--no-sandbox', '--disable-setuid-sandbox']});
    const page = await browser.newPage();
    await page.setViewport({
        width: width,
        height: height,
        deviceScaleFactor: 1,
    });            
    await page.goto(url, { waitUntil: 'networkidle2' });
    const buffer = await page.screenshot({encoding: 'binary'});
    await browser.close();

    return buffer
}

module.exports = {
    getScreenshot
}