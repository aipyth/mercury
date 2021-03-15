const gulp = require('gulp')
const pug = require('gulp-pug')
const plumber = require('gulp-plumber')
const pugLinter = require('gulp-pug-linter')
const bemValidator = require('gulp-html-bem-validator')
const fs = require('fs')

// const writeStatic = function(){
//     fs.readFile('src/templates/index.html', 'utf8', (err, data)=>{
//         if (err) console.log( err)
//         fs.writeFile('src/templates/index.html', ("{% load static %}" + data), (err)=>{
//             if (err) console.log( err)
//         })
//         console.log("Load static written sucess")
//     })
// }
    


module.exports = function pug2html(cb) {
    return gulp.src('front/pages/*.pug')
    .pipe(plumber())
    .pipe(pugLinter({reporter: 'default'}))
    .pipe(pug())
    .pipe(bemValidator())
    .pipe(gulp.dest('src/api/templates'))
}