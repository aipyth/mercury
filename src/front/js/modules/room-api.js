
class Room {
    constructor(id) {
        return api.roomsGet(id).then(data => {
                this.id = data.id
                this.time_schema = data.time_schema
                this.subjects = data.subjects
                this.period = data.period
                this._subects_list = this.subjects.map(x => x.name)
            }).catch(err => console.error(err))
    }

    addSubject(week_id, day, subj_name, order) {
        // this.weeks[week_id][day]['subjects'][order] = subj_name + ':'
        if(this._subjects_list.includes(subj_name)){
            this.subjects[subj_name].days_and_orders[day] = [order]
           return api.subjectsUpdate(this.subjects[subj_name].id, subj_name, this.id, this.subjects[subj_name].days_and_orders)
                .then(res => api.subjectsGet())
                .catch(err => console.error(`Error updating subject: ${err}`))
                .then(data => {this.subjects = data})
                .catch(err => console.error(`Error getting subject: ${err}`))
        } else {
            this.subjects[subj_name] = {days_and_orders: {[day]: [order]}}
            this._subects_list.push(subj_name)
            return api.subjectsCreate(subj_name, this.id, subjects[subj_name].days_and_orders)
                .then(res => api.subjectsGet())
                .catch(err => console.error(`Error posting subject: ${err}`))
                .then(data => {this.subjects = data})
                .catch(err => console.error(`Error getting subject: ${err}`))
        } 
    }

    display(daytag) {
        const displaySubject = (day, order, subject) => {
                var subjsection = day == 0 ? 
                    $('.shedule-main').find(`div.subj-section.${order+1}`) :
                    $(`.shedule.${day}`).find(`div.subj-section.${order+1}`)
                subjsection.find('.name').html(subject.name)
        }
        const displayScheme = (el, day) => {
            el.find('h2').html(day)
            for (let i = 1; i <= 5; i++){
                el.find(`div.subj-section.${i}`).find('.time').html(this.time_schema[i])
            }
        }
        var currentDay = new Date().getDay() - 1
        if (currentDay == -1) currentDay++
        const arrayOfWeekdays = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
        if(daytag) currentDay = arrayOfWeekdays.indexOf(daytag)
        arrayOfWeekdays = arrayOfWeekdays.slice(currentDay-1) + arrayOfWeekdays.slice(0, currentDay)
        displayScheme($('.shedule-main'), arrayOfWeekdays[currentDay])
        for (let i = 1; i <= 5; i++) {
            displayScheme($(`.shedule.${i}`), arrayOfWeekdays[i])
        }
        for(let subject of this.subjects){
            for(day in subject.days_and_orders){
                for(order of subject.days_and_orders[day]){
                    displaySubject(arrayOfWeekdays.indexOf(day), order, subject)
                }
            }
        }
    }

}

Room.create = api.roomsCreate

module.exports = Room;