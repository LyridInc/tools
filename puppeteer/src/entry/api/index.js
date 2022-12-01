
const router = require('express').Router();
const sharp = require('sharp')
const puppeteer = require('../../puppeteer')

router.get('/screenshot/:width/:height?', async (req, res, next) => {
    const baseUrl = req.query.url
    const width = parseInt(req.params.width, 10)
    const height = parseInt(req.params.height, 10)
    const transformWidth = parseInt(req.query.width, 10) || null

    try {
        
        let buffer = await puppeteer.getScreenshot(baseUrl, width, height)

        if (transformWidth != null) {
            const transformer = sharp(buffer)
            transformer.resize(transformwidth, height, {
                fit: 'inside',
                withoutEnlargement: true,
              })
            
            buffer = transformer.toFormat('png').toBuffer()
        }
        

        res.type(`image/png`).status(200).send(buffer)
    } catch (error) {
        if (error.statusCode === 404) {
            next()
            return
        }
        next(error)
    }
})

module.exports = router