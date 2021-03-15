const gulp = require('gulp')
const imagemin = require("gulp-imagemin")

module.exports = function pug2html(cb) {
    return gulp.src('front/img/*')
    .pipe(imagemin([
        imagemin.gifsicle({interlaced: true}),
        imagemin.mozjpeg({quality: 75, progressive: true}),
        imagemin.optipng({optimizationLevel: 5}),
        imagemin.svgo({
            plugins: [
                {removeViewBox: true},
                {cleanupIDs: false}
            ]
        })
    ]))
    .pipe(gulp.dest('src/api/schedule/static/schedule/img'))

}