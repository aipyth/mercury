const catchError = (err) => console.error(new Error(err))

const resolve = (response, status) => new Promise((resolve, reject) => {
    if(response.status == status) response.json().then(data => resolve(data))
    else response.json().then(data => reject(data))
})

var token = utilits.read_cookie('token')

const api = {
    timeSchemesGet: id => {
    id = id ? id : ''
    return window.fetch(`/api/timeschemes/${id}`, 
    {
        headers: 
        {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'GET' 
    }).then(response => resolve(response, 200))
    .catch(err => catchError(err))
    },

    timeSchemesCreate: (name, items, isPublic = true) => window.fetch(`/api/timeschemes`, 
        {
            headers: 
            {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8',
                'Authorization': `Token ${token}`
            },
            method: 'POST' ,
            body: JSON.stringify({
                name: name,
                items: items,
                public: isPublic
            })
        }).then(response => resolve(response, 201))
        .catch(err => catchError(err)),

    timeSchemesUpdate: (id, name, items, isPublic = true) => window.fetch(`/api/timeschemes${id}`, 
    {
        headers: 
        {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'PUT' ,
        body: JSON.stringify({
            name: name,
            items: items,
            public: isPublic
        })
    }).then(response => resolve(response, 201))
    .catch(err => catchError(err)),

    timeSchemesDelete: id => window.fetch(`/api/timeschemes/${id}`, 
    {
        headers: 
        {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'DELETE' 
    }).then(response => resolve(response, 200))
    .catch(err => catchError(err)),

    roomsGet: id => window.fetch(`/api/rooms/${id ? id : ''}`,
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'GET'
    }).then(response => resolve(response, 200))
    .catch(err => catchError(err)),

    roomsCreate: room_settings => window.fetch(`api/rooms/`, 
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'POST',
        body: JSON.stringify(room_settings)
    }).then(response => resolve(response, 201))
    .catch(err => catchError(err)),

    roomsUpdate: room_settings => window.fetch(`api/rooms/`, 
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'PUT',
        body: JSON.stringify(room_settings)
    }).then(response => resolve(response, 201))
    .catch(err => catchError(err)),

    roomsDelete: id => window.fetch(`/api/rooms/${id}`,
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'DELETE'
    }).then(response => resolve(response, 200))
    .catch(err => catchError(err)),

    subjectsGet: id => window.fetch(`/api/subjects/${id}`,
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'GET'
    }).then(response => resolve(response, 200))
    .catch(err => catchError(err)),

    subjectsCreate: (name, room, days_and_orders) => window.fetch(`/api/subjects/`,
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'POST',
        body: JSON.stringify({
            days_and_orders: days_and_orders,
            room: room,
            name: name
        })
    }).then(response => resolve(response, 201))
    .catch(err => catchError(err)),

    subjectsUpdate: (id, name, room, days_and_orders) => window.fetch(`/api/subjects/${id}`,
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'PUT',
        body: JSON.stringify({
            days_and_orders: days_and_orders,
            room: room,
            name: name
        })
    }).then(response => resolve(response, 201))
    .catch(err => catchError(err)),

    subjectsDelete: id => window.fetch(`/api/subjects/${id}`,
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
            'Authorization': `Token ${token}`
        },
        method: 'DELETE'
    }).then(response => resolve(response, 200))
    .catch(err => catchError(err)),

    createUser: (email, password) => window.fetch(`/api/create-user/`,
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
        },
        method: 'POST',
        body: JSON.stringify({
            email: email,
            password: password
        })
    }).then(response => resolve(response, 201)),

    createToken: (username, password) => window.fetch(`api/api-token-auth/`,
    {
        headers: {
            'Accept': 'application/json, text/plain',
            'Content-Type': 'application/json;charset=UTF-8',
        },
        method: 'POST',
        body: JSON.stringify({
            username: username,
            password: password
        })
    }).then(response => resolve(response, 200)),

    isSigned: () => token != null
}

module.exports = api;