const fs = require('fs')

const {promisify} = require('util')

const readFile = promisify(fs.readFile)

const readDir = promisify(fs.readdir)

const writeFile = promisify(fs.writeFile)

var path = './front/js/modules'

Array.prototype.forEachAsync = async function (callback) {
    for (let i =0; i < this.length; i++) {
       await callback(this[i], i, this)
    }
}

const getSubtring = (str, from, to) => {
    let length = to - from
    return str.substr(from, length + 1)

}

const getExport = file_data => {
    let first_index = file_data.indexOf('module.exports')
    let length = file_data.indexOf('\n', first_index) != -1 ? 
        file_data.indexOf('\n', first_index) - first_index : file_data.indexOf(';', first_index) - first_index
    if (length < 0) throw new Error(`Error, plese put ';' at the end of module.exports = ...`)
    let substring = file_data.substr(first_index, length + 2)
    return substring.split('= ')[1].slice(0, -1)
}

const clearExport = file_data => {
    let first_index = file_data.indexOf('module.exports')
    let length = file_data.indexOf('\n', first_index) != -1 ? 
        file_data.indexOf('\n', first_index) - first_index : file_data.indexOf(';', first_index) - first_index
    if (length < 0) throw new Error(`Error, plese put ';' at the end of module.exports = ...`)
    let substring = file_data.substr(first_index, length + 2)
    return file_data.replace(substring, '')
}

const stringIterator = (str, substr, callback) => {
    var index = str.indexOf(substr)
    while (index != -1) {
        str = callback(str, index)
        index = str.indexOf(substr)
    }
}

const collectRequires = file_data => {
    var requires = {}
    stringIterator(file_data, 'require', (data, index) => {
        let first_index = data.lastIndexOf('const', index)
        // let length = data.indexOf('\n', first_index) - first_index
        // let substring = data.substr(first_index, length + 1)
        let substring = getSubtring(file_data, first_index, data.indexOf('\n', first_index))

        data = data.replace(substring, '')

        let req_from = substring.indexOf("'") + 1
        let req_to = substring.indexOf(')') - 2
        let req = getSubtring(substring, req_from, req_to).split('/')[1] + '.js'

        let varname_from = substring.indexOf(' ') + 1
        let varname_to = substring.indexOf('=', varname_from) - 2
        let varname = getSubtring(substring, varname_from, varname_to)
        requires[req] = varname
        file_data = data
        return data
    })
    return [file_data, requires]
}

const addRequiredModule = (require, exported) => require != exported ? `var ${require} = ${exported};\n` : ''


const compileFiles = async (path_to) => {
    var array_of_files_data = [] 
    var required_vars = []
    var main_path = path + '/../main_mod.js'
    var data_buffer = await readFile(main_path)
    let [main_data, requires] = collectRequires(data_buffer.toString())

    const files = await readDir(path)
    await Object.keys(requires).forEachAsync(async filename => {
        if (files.includes(filename)) {
            var data = await readFile(path + '/' + filename)
            required_vars.push(addRequiredModule(requires[filename], getExport(data.toString())))
            data = clearExport(data.toString())
            array_of_files_data.push(data + '\n')
        }
        else console.error(`Error, file '${filename}' does not exist by path ${path}, so require('${filename}') has been ignored`)
        
    })
    main_data = array_of_files_data.join('') + required_vars.join('') + main_data
    await writeFile(path_to, main_data)
    console.log('Compiled modules into one file succesfully!')
}


compile = cb => {
    compileFiles('./front/js/main.js')
    cb()
}

module.exports = compile
