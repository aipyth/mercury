function getAuth(){
    const email = $('#email').val()
    const password = $('#password').val()
    // console.log(password)
    return [email, password]
}  

const auth = {
    username: '',

    menu: {
        el: $('.modal'),
        button: $('.modal #next'),
        log: () => {
            $('section.modal').addClass('vis')
            $('section.modal .sign-up p').show()
            $('.sign-head h2').text(`Log in`)
            $('#next h4').text(`Log in`)
        },
        sign: () => {
            $('section.modal').addClass('vis')
            $('section.modal .sign-up p').hide()
            $('.sign-head h2').text(`Sign up`)
            $('#next h4').text(`Sign up`)
        },
        hide: function() {this.el.removeClass('vis')}
    },

    signUp: () => {
        var sign_data = getAuth()
        return api.createUser(sign_data[0], sign_data[1])
    },

    logIn: function() {
        var sign_data = getAuth()
        this.username = sign_data[0]
        return api.createToken(sign_data[0], sign_data[1])
    }
}

module.exports = auth;