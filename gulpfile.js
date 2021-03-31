const gulp = require('gulp')

const pug2html = require('./gulp/tasks/pug2html')
const styles = require('./gulp/tasks/styles')
const script = require('./gulp/tasks/script')
const { series, task } = require('gulp')
const { stream } = require('browser-sync')
const imageMinifier = require('./gulp/tasks/image-minifier')
const server = require('browser-sync').create()
const clean = require('./gulp/tasks/clean')
const compileFiles = require('./gulp/tasks/compiler')
const fs = require('fs');

const writeStatic = (cb) => {
    fs.readFile('src/api/templates/index.html', 'utf8', (err, data)=>{
        if (err) console.log( err)
        fs.writeFile('src/api/templates/index.html', ("{% load static %}" + data), (err)=>{
            if (err) console.log( err)
        })
        console.log("Load static written sucess")
    })
    cb();
}

function serve(cb) {
    gulp.watch('front/css/**/*.sass', series(styles)).on('change', server.reload)
    gulp.watch('front/js/modules/*.js', series(compileFiles))
    gulp.watch('front/js/main_mod.js', series(compileFiles))
    gulp.watch('front/js/main.js', series(script))
    gulp.watch('front/pages/**/*.pug', series(pug2html, writeStatic))
    gulp.watch('src/api/templates/*.html').on('change', server.reload)
    // gulp.watch('front/img/*', series(imageMinifier)).on('change', server.reload)
    return cb
}  
module.exports.build = series(clean, pug2html, writeStatic, styles, compileFiles, script)
module.exports.start = series(module.exports.build, serve)


