const utilits = {
    read_cookie: (key, skips) => {
        if (skips == null)
                skips = 0;
        var cookie_string = "" + document.cookie;
        var cookie_array = cookie_string.split("; ");
        for (var i = 0; i < cookie_array.length; ++i)
        {
                var single_cookie = cookie_array[i].split("=");
                if (single_cookie.length != 2)
                    continue;
                var name  = single_cookie[0];
                var value = single_cookie[1];
                if (key == name && skips -- == 0)
                        return value;
        }
        return null;
    },

    write_cookie: async (name, value) => {
            var expiration_date = new Date();
            expiration_date.setYear(2025);
            expiration_date = expiration_date.toGMTString();
            var cookie_string = name + "=" + value + "; expires=" + expiration_date;
            document.cookie = cookie_string;
    },

    saveToken: async (data) => {
        token = data.token
        await write_cookie('token', token)
    }
}
module.exports = utilits;