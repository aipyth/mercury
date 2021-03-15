const gulp = require('gulp')
const terser = require('gulp-terser')
const sourcemaps = require('gulp-sourcemaps')
const babel = require('gulp-babel')
const rename = require('gulp-rename')
const eslint = require('gulp-eslint') 

module.exports = function script(cb) {
    return gulp.src('front/js/*.js')
    .pipe(babel(
        {
            presets: [
                [
                  '@babel/preset-env',
                  {
                    targets: {
                      esmodules: true,
                    },
                  },
                ],
              ],
        }))
    .pipe(eslint({
        extends: ["standart", "htmlacademy/es6"],
        globals: [
            'jQuery',
            '$',
            'this',
        ],
    }))
    .pipe(eslint.format())
    // .pipe(sourcemaps.init())
    //     .pipe(terser())
    // .pipe(sourcemaps.write())
    .pipe(rename({suffix: '.min'}))
    .pipe(gulp.dest('src/api/schedule/static/schedule/js'))
}