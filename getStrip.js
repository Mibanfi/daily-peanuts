import puppeteer from 'puppeteer';

const browser = await puppeteer.launch();
const page = await browser.newPage();

page
    .waitForSelector('.Comic_comic__image_isStrip__eCtT2')
    .then(async ()=>{
    var imgsrc = await page.evaluate(()=>{
        return document.getElementsByClassName('Comic_comic__image__6e_Fw Comic_comic__image_isStrip__eCtT2').item(0).src;
    })
    console.log(imgsrc);
    browser.close();
})

await page.goto("https://www.gocomics.com/peanuts");