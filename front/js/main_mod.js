const {read_cookie, write_cookie, saveToken} = require('modules/utilits')
const api = require('modules/server-api')
const auth = require('modules/auth')
const Room = require('modules/room-api')

///////////////// SOME FUNCTIONS //////////////////

const goTo = (element) => {
    var offset =  element.offset()
    offset.left -= 20;
    offset.top -= 50;
    $('html, body').animate({
        scrollTop: offset.top,
        scrollLeft: offset.left
    });
}

//////////////////////////////////////////////////////

////////////////////// NOT AUTH //////////////////////

$('.auth-inner .sign-underline#log').click(() => {
   auth.menu.log()
})

$('.auth-inner .sign-underline#sign').click(() => {
   auth.menu.sign()
})

//////////////////////////////////////////////////

/////////////// MENU JS CODE /////////////////////

var this_week = 0

$('#sign').on('click', ()=>{
    $('section.modal').addClass('vis')
})

$('#shemes').click(()=>{
    $('section.content').addClass('hidden')
    $('.room-settings').removeClass('hidden')
    displayRooms()
})

$(`.header .logo img`).click(() => {
    if(picked_room_id){
        $('section.content').removeClass('hidden')
        $('.room-settings').addClass('hidden')
    } else {
        $('.room-settings').addClass('hidden')
        selectRoomMessage.show()
    }
})

$('#nextweek').click(() => {
    const wek = $('.nextweek-button h3')
    wek.html('Week 2')
    $('#nextweek').addClass('thisweekhiden')
    $('#prevweek').removeClass('thisweekhiden')
    this_week = 1
    displayDays(Weeks[this_week])
})

$('#prevweek').click(() => {
    $('.nextweek-button h3').html('Week 1')
    $('#prevweek').addClass('thisweekhiden')
    $('#nextweek').removeClass('thisweekhiden')
    this_week = 0
    displayDays(Weeks[this_week])
})


//////////////////////////////////////////////////////

/////////////////// AUTHORIZATION ///////////////////

const checkForAuthorization = function() {
    if (api.isSigned()) {
        $('.header .menu').removeClass('hidden')
        $('section.auth-enter').addClass('hidden')
    } else {
        $('.header .menu').addClass('hidden')
        $('section.content').addClass('hidden')
        $('section.auth-enter').removeClass('hidden')
    }
    return api.isSigned()
}

function signUpFail(data) {
    
    if (data.email) $('#sign-errors').html(`<p> *Error ${data.email} </p>`)
    else  $('#sign-errors').html(`<p> *Error ${data.detail} </p>`)
    setTimeout(() => {
        $('#sign-errors').html(``)
    }, 7000);
}

async function logInSucessful(data, username) {
    $('#sign-sucess').html(`<h2>Welcome to the club, buddy </h2>
    <h2>${username}</h2>`)
    $('#sign-form').addClass('hidden')
    $('#sign-sucess').removeClass('hidden')
    setTimeout(() => {
        $('.modal').removeClass('vis')
    }, 5000);
    await saveToken(data)
    checkForAuthorization()
}

function logInFail(data) {
    if (data.detail) $('#sign-errors').html(`<p> *Error ${data.detail} </p>`)
    else  $('#sign-errors').html(`<p> *Error ${data.non_field_errors[0]} </p>`)
    setTimeout(() => {
        $('#sign-errors').html(``)
    }, 7000);
}

auth.menu.el.find('.cross').on('click', ()=> { auth.menu.hide() })

$(document).mouseup((e)=>{
    var modal = auth.menu.el.find('.sign-block')
    if(!modal.is(e.target) && modal.has(e.target).length === 0){
        auth.menu.hide()
    }
})

auth.menu.button.on('click', ()=>{
   if(auth.menu.button.find('h4').text() == 'Sign up'){
       auth.signUp()
        .then(
            result => auth.logIn()
            .then(data => logInSucessful(data, auth.username))
            )
        .catch(fail => signUpFail(fail))
   }
   else {
       auth.logIn()
        .then(data => logInSucessful(data, auth.username))
        .catch(fail => logInFail(fail))
   }
})

$('#sign-underline').on('click', ()=>{
    auth.menu.sign()
})
///////////////////////////////////////////////////////////

//////////////////// ROOM MANAGE ///////////////////

var picked_room_id = parseInt(read_cookie('room'))
var room = {}

const checkForRoom = async function() {
    if(token) {
        if(picked_room_id) {
            room = await new Room(picked_room_id)
            room.display()
        } else {
            selectRoomMessage.show()
        }
    } 
}

const checkForUrl = async function () {
    var url_slug = window.location.href.split('/')[3]
    if (url_slug) {
        var rooms = await api.roomsGet()
        rooms = rooms.results 
        var slug_correct = false
        for (var room of rooms) {
            if (room.slug == url_slug) {
                await write_cookie('room', room.id)
                slug_correct = true
            } 
        }
        if(!slug_correct) alert("Can't get room by this url, is it correct?")
    }
}


$('.shedule-main').find('.subj-section .name').on('focusout', function() {
    const day = $('.shedule-main').find('h2').text()
    const subj_name = $(this).text()
    const order = $(this).parent().children().index($(this))
    room.addSubject(week_id, day, subj_name, order)
    console.log(Weeks)
})

$('.shedule').on('click', function(){
    const daytag = $(this).find('h2').text()
    room.display(daytag)
    // displayDays(Weeks[this_week], getDayId(daytag))
    
})


var selectRoomMessage = {
    show: function() {
        $('section.content').addClass('hidden')
        $('.select-room-message').removeClass('hidden')
        $('#room-sett').click(()=>{
            $('.select-room-message').addClass('hidden')
            $('.room-settings').removeClass('hidden')
            displayRooms()
        })
    },
    hide: function() {
        $('section.content').removeClass('hidden')
        $('.select-room-message').addClass('hidden')
    }
}


$(async () => {
    if (token) await checkForUrl()
    checkForAuthorization() ? 
    await checkForRoom() : null
})

///////////////////////////////////////////////////////

///////////////////// SHEDULE //////////////////////


// var Weeks= [{
//     Monday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Tuesday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Wednesday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Thursday: {
//                 subjects:[':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Friday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Saturday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
// }, {
//     Monday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Tuesday:{
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Wednesday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Thursday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Friday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
//     Saturday: {
//                 subjects: [':', ':', ':', ':', ':'],
//                 times: ['8:30 - 10:05', '10:25 - 12:00', '12:20 - 13:55', '14:15 - 15:50', '16:10 - 17:45'],
//                 tasks: []
//             },
// }]

// var all_subjects_server = []

// const postSubjects = () => {
//     var array_of_all_subjects = []
//     var object_of_days_and_orders = {}
//     for (let day of ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]) {
//         for(let week of Weeks){
//             var oneday_subjects = week[day].subjects
//             for (let item of oneday_subjects){
//                 item = item.split(':')[0]
//                 if (item == '') continue
//                 if (item in object_of_days_and_orders){
//                     var obj = object_of_days_and_orders[item]
//                     if(obj[day].includes(item)) obj[day].push(oneday_subjects.indexOf(item + ':'))
//                     else obj[day] = [oneday_subjects.indexOf(item + ':')]
//                     object_of_days_and_orders[item] = obj
//                 } else {
//                     var obj = {}
//                     obj[day] = [oneday_subjects.indexOf(item+':')]
//                     object_of_days_and_orders[item] = obj
//                 }
//             }
//             array_of_all_subjects = array_of_all_subjects.concat(week[day].subjects)
//         }
//     }
//     var all_subjects = new Set(array_of_all_subjects)
//     all_subjects = Array.from(all_subjects)
//     for (let item of all_subjects) all_subjects[all_subjects.indexOf(item)] = item.split(':')[0]
//     all_subjects = new Set(all_subjects)
//     all_subjects.delete('')

//     var server_subject_names = (() => {
//         let arr = []
//         for (let subj of all_subjects_server){
//             arr.push(subj.name)
//         }
//         return arr
//     })()

//     for(var subject of all_subjects){
//         if (server_subject_names.includes(subject)){
//             // window.fetch('/api/subjects/',  {
//             //     headers: {
//             //         'Accept': 'application/json, text/plain',
//             //         'Content-Type': 'application/json;charset=UTF-8',
//             //         'Authorization': `Token ${token}`
//             //     },
//             //     method: 'PUT', 
//             //     body: JSON.stringify({

//             //         days_and_orders: object_of_days_and_orders[subject],
//             //         name: subject,
//             //         room: picked_room_id
//             //     })
//             // })
//             // .then (
//             //     async response => await responseResult(response, 201, log, () => {throw Error(response)})
//             // )
//             api.subjectsUpdate(subject, picked_room_id, object_of_days_and_orders[subject])
//         } else {
//             // window.fetch('/api/subjects/',  {
//             //     headers: {
//             //         'Accept': 'application/json, text/plain',
//             //         'Content-Type': 'application/json;charset=UTF-8',
//             //         'Authorization': `Token ${token}`
//             //     },
//             //     method: 'POST', 
//             //     body: JSON.stringify({

//             //         days_and_orders: object_of_days_and_orders[subject],
//             //         name: subject,
//             //         room: picked_room_id
//             //     })
//             // })
//             // .then (
//             //     async response => await responseResult(response, 201, log, () => {throw Error(response)})
//             // )
//             api.subjectsCreate(subject, picked_room_id, object_of_days_and_orders[subject])
//         }    
//     }
// }
//     let response = await window.fetch('/api/rooms/', 
//     {
//         headers: {
//             'Accept': 'application/json, text/plain',
//             'Content-Type': 'application/json;charset=UTF-8',
//             'Authorization': `Token ${token}`
//         },
//         method: 'GET' 
//     })
//     let rooms_list = await response.json()
//     for (let room of rooms_list.results) {
//         if (room.id == id){
//             return room
//         }
//     }
// }

// const getSchema = api.timeSchemesGet
//     let response = await window.fetch(`/api/timeschemes/${id}`, 
//     {
//         headers: 
//         {
//             'Accept': 'application/json, text/plain',
//             'Content-Type': 'application/json;charset=UTF-8',
//             'Authorization': `Token ${token}`
//         },
//         method: 'GET' 
//     })
//     var scheme = await response.json()
//     return scheme
// }
// const getWeek = async () => {
//     selectRoomMessage.hide()
//    var room = await api.roomsGet(picked_room_id)
//    var timeschemes = await api.timeSchemesGet(room.time_schema) 
//    all_subjects_server = room.subjects
//    var array_of_times = []
//    for (let item in timeschemes.items) {
//        var time = '' + timeschemes.items[item].Start + ' - ' + timeschemes.items[item].Stop
//        array_of_times.push(time)
//    }

//    var shortperiod = (room.period == 7)
//    if (shortperiod) $('section.content .nextweek-field').addClass('hidden')

//    // Adding timeschemes from server to client var Weeks
//    for (let day of ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]) {
//        Weeks[0][day].times = array_of_times
//        if (!shortperiod) Weeks[1][day].times = array_of_times
//    }
//    // Adding subjects from server to client var Weeks
//    for (let subject of all_subjects_server){
//        for(let day in subject.days_and_orders){
//             for (let order of subject.days_and_orders[day]){
//                 Weeks[0][day].subjects[order] = subject.name
//                 if (!shortperiod) Weeks[1][day].subjects[order] = subject.name
//             }
//         }
//     }

// }

// const getDayId = (daytag) => ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"].indexOf(daytag)


// const displayShedule = (jquerObj, week_day) => {
//     for (let index = 0; index < 5; index++) {
//         var subjsection = jquerObj.find(`div.subj-section.${index+1}`)
//         subjsection.find('.name').html(week_day['subjects'][index].split(':')[0])
//         subjsection.find('.time').html(week_day['times'][index])
//     }
// } 

// const displayDays = (week, mainDay = -1) => {
//     const arrayOfWeekdays = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
//     var currentDay = mainDay == -1 ? (new Date().getDay() - 1) : mainDay
//     if (mainDay == -1) mainDay++
//     $('.shedule-main').find('h2').html(arrayOfWeekdays[currentDay])
//     displayShedule($('.shedule-main'), week[arrayOfWeekdays[currentDay]])
//     for (let i = 1, j = 0; i <= 5; i++, j++) {
//         if (j == currentDay) {
//             i--
//             continue
//         }
//         $(`.shedule.${i}`).find('h2').html(arrayOfWeekdays[j])
//         displayShedule($(`.shedule.${i}`), week[arrayOfWeekdays[j]])
//     }
// }

// const changeSheduleSubject = (thisweek) => {
//     const day = $('.shedule-main').find('h2').text()
//     const parsefrom = $('.shedule-main .subjects .subj-section')
//     for (let i = 0; i < 5; i++) {
//         Weeks[thisweek][day]['subjects'][i] = $(parsefrom[i]).find('.name').text() + ':'
//     }
//     postSubjects()
// }



//////////////////////////////////////////////////////

///////////////////// ROOOM SETTINGS //////////////////////

var number_of_settings = $('.settings-content').children().length
var setting_id = 0

$('#settings-left').click(() => {
    $($('.settings-content').children()[setting_id]).addClass('hidden')
    setting_id = (number_of_settings + setting_id - 1) % number_of_settings
    $($('.settings-content').children()[setting_id]).removeClass('hidden')
})

$('#settings-right').click(() => {
    $($('.settings-content').children()[setting_id]).addClass('hidden')
    setting_id = (setting_id + 1) % number_of_settings
    $($('.settings-content').children()[setting_id]).removeClass('hidden')
})

///////////////// ROOM CREATION ///////////////////////

var Time_sheme = {
    1: {
        Start: '8:30',
        Stop: '10:05'
    },
    2: {
        Start: '10:25',
        Stop: '12:00'
    },
    3: {
        Start: '12:20',
        Stop: '13:55'
    },
    4: {
        Start: '14:15',
        Stop: '15:50'
    },
    5: {
        Start: '16:10',
        Stop: '17:45'
    },
}

var Time_shemes_list = [];

const createSlider = (id, min, max, startVal, secVal) => {
    var slider = {
    id: ('#' + id),
    getSilderData: function() {
        var inp1 = $(this.id).find('#a')
        var inp2 = $(this.id).find('#b')
        this.val1 = parseInt(inp1.val()) < parseInt(inp2.val()) ? inp1.val() : inp2.val()
        this.val2 = parseInt(inp1.val()) > parseInt(inp2.val()) ? inp1.val() : inp2.val()
    },
    oninput: function (callback) {
            var id_targert = '#' + $(this.id).attr('for')
            this.getSilderData()
            callback($(id_targert))
            addEventListener('input', e => {
                this.getSilderData()
                callback($(id_targert))
            }, false);
            addEventListener('change', e => {
                this.getSilderData()
                callback($(id_targert))
            });
        }
    }
        $(slider.id).html(`
        <div class='wrap' role='group' aria-labelledby='multi-lbl'>
            <label class='sr-only' for='a'></label>
            <input id='a' type='range' max='${max}' value='${startVal}' min='${min}'>
            <label class='sr-only' for='b'></label>
            <input id='b' type='range' max='${max}' value='${secVal}' min='${min}'>
        </div>
        `)
        return slider
}


const changeSheme = (new_sheme) => {
    Time_sheme = new_sheme.items
    Time_sheme.id = new_sheme.id
    $('.sheme .slider-range').each(function(i) {
        var index = i + 1
        var value1 = Time_sheme[index].Start;
        var value2 = Time_sheme[index].Stop;
        const count_time = value => {
            value = value.split(':')
            for (let j=0; j < 2; j++) value[j] = parseInt(value[j])
            value[0] = (value[0] - 6) * 12
            value[1] = (value[1]/5)
            value = value[1] + value[0]
            return value
        }
        value1 = count_time(value1)
        value2 = count_time(value2)
        $(this).find('#a').val(value1)
        $(this).find('#b').val(value2)
    })
}


$('.sheme .slider-range').each(function(i) {
    var index = i + 1
    var slider = createSlider(`slider-range${index}`, 0, 216, 54, 96)
    slider.oninput(function(target){
        var values = [slider.val1, slider.val2]
        for (let j = 0; j < 2; j++) {
            var hours = Math.floor(values[j] / 12)
            hours = hours + 6
            hours = (hours + '').split('')
            if (hours.length == 1) hours.unshift('0')
            var minutes = values[j] % 12
            minutes = minutes * 5
            minutes = (minutes + '').split('')
            if (minutes.length == 1) minutes.unshift('0')
            var sel = ['.Start', '.Stop'][j]

            Time_sheme[index][sel.split('.')[1]] = hours[0] + hours[1] + ':' + minutes[0] + minutes[1]

            var hs = target.find(sel).find('.hours')
            var mins = target.find(sel).find('.minutes')
            for (let k = 0; k < 2; k++) {
                hs.find(`.timenumber:eq(${k}) span`).text(hours[k])
                mins.find(`.timenumber:eq(${k}) span`).text(minutes[k])
            }

        }
    })
})

// const timeSchemesPost = async (sheme_name) {
    
    // window.fetch('/api/timeschemes/', 
    // {
    //     headers: {
    //         'Accept': 'application/json, text/plain',
    //         'Content-Type': 'application/json;charset=UTF-8',
    //         'Authorization': `Token ${token}`
    //     },
    //     method: 'POST', 
    //     body: JSON.stringify({
    //         name: sheme_name,
    //         items: Time_sheme
    //     })
    // })
    // .then(
    //     response => responseResult(response, 201, ((response)=>{
    //         response.json().then(data => log(data))
    //     }), ()=>{throw Error(`Error response status is ${response.status}`)})
    // )
    // .catch((err) => log(err))
// }

// async function createRoomPost(room_settings) {
    
//     window.fetch('/api/rooms/', 
//     {
//         headers: {
//             'Accept': 'application/json, text/plain',
//             'Content-Type': 'application/json;charset=UTF-8',
//             'Authorization': `Token ${token}`
//         },
//         method: 'POST', 
//         body: JSON.stringify(room_settings)
//     })
//     .then(
//         response => responseResult(response, 201, data => log(data), data =>{throw Error(`Error response status is ${data}`)})
//     )
//     .catch((err) => log(err))
// }

// async function getTimeShemes() {
    
//     let response = await window.fetch('/api/timeschemes/', 
//     {
//         headers: {
//             'Accept': 'application/json, text/plain',
//             'Content-Type': 'application/json;charset=UTF-8',
//             'Authorization': `Token ${token}`
//         },
//         method: 'GET'
//     })
//    return await response.json()
// }

$('#shemes').click(async function () {
    var time_shemes = await api.timeSchemesGet()
    Time_shemes_list = time_shemes
    var options = `<option> ... </option>`
    for (let sheme of time_shemes.results) {
        options = options + `<option>${sheme.name}</option>`
    }
    $('#namesheme').html(options)
})


$('#submitsheme').click(async () => {
    var sheme_name = $('.label #timeshemename').val()
    if (sheme_name != '')
    {
        await api.timeSchemesCreate(sheme_name, Time_sheme);
        var time_shemes = await api.timeSchemesGet()
        var options = `<option> ... </option>`
        for (let sheme of time_shemes.results) {
            if (sheme.name == sheme_name) options = options + `<option selected>${sheme.name}</option>`
            else options = options + `<option>${sheme.name}</option>`
        }
        $('#namesheme').html(options)
    } else {
        $('.label #timeshemename').attr('placeholder', "Input name please")
        goTo($('.label #timeshemename'))
    }

})

$('#createroom').click(async () => {
    if ($('#namesheme').val() == '...') {
        $('#error-message h4').text(`Error Time sheme field is empty`)
        return
    }
    for(let inp_el of ['start-date', 'end-date', 'name' ]) {
        let inp_el_formated = inp_el.charAt(0).toUpperCase() + inp_el.slice(1)
        if($(`.label #${inp_el}`).val() == '') {
            inp_el_formated = inp_el_formated.split('-')
            inp_el_formated[1] =  inp_el_formated[1] == undefined ? '' : inp_el_formated[1]
            $('#error-message h4').text(`Error ${inp_el_formated[0] + ' ' + inp_el_formated[1]} field is empty`)
            return
        }
    }
    var room_settings = {}
    room_settings.addField = (field_name, id=field_name) => {
        room_settings[field_name] = $(`.label #${id}`).val()
    }
    room_settings.addField('name')
    room_settings.period = parseInt($(`.label #period`).val())
    room_settings.addField('start_date', 'start-date')
    room_settings.addField('end_date', 'end-date')
    room_settings.addField('slug')
    room_settings.public = ($(`.label #public`).val() == 'on')
    room_settings['time_schema'] = Time_sheme.id
    await api.roomsCreate(room_settings)
    displayRooms()
    $($('.settings-content').children()[setting_id]).addClass('hidden')
    setting_id = (setting_id + 1) % number_of_settings
    $($('.settings-content').children()[setting_id]).removeClass('hidden')
})  

$('#namesheme').change(function() {
    var selected_shema = (() => { 
        for(let shema of Time_shemes_list.results) {
            if (shema.name == $(this).val()) return shema
        }
    })();
    changeSheme(selected_shema)
})

///////////////////// PICK ROOM ///////////////////////////

const getShemeName = async id => {
    var data = await api.timeSchemesGet(id)
    return data.name
}

const displayRooms = async () => {
    var list_of_rooms = await api.roomsGet()
    var room_list_html = $('.settings-content').children()[1]
    $(room_list_html).html('')
    for (let room of list_of_rooms.results)
        $(room_list_html).append(`
            <div class="room-block">
                <div class="room-name">Name: <span class="room-info">${room.name}</span></div>
                <div class="room-id">Id: <span class="room-info">${room.id}</span></div>
                <div class="room-sheme">Time scheme: <span class="room-info">${await getShemeName(room.time_schema)}</span></div>
            </div>
        `)
    var childs = $($('.settings-content').children()[1]).children()
    $(childs).each(function() {
        $(this).on('click', async function() {
            let id = $(this).find('.room-id').find('.room-info').text()
            picked_room_id = parseInt(id)
            await write_cookie('room', picked_room_id)
            await getWeek()
            displayDays(Weeks[this_week])
            $('section.content').removeClass('hidden')
            $('.room-settings').addClass('hidden')
        })
    })
}



