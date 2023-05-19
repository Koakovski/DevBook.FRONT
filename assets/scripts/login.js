$("#login-form").on("submit", login);

function login(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    $.ajax({
        url: "api/login",
        method: "POST",
        data,
    })
        .done(successHandler)
        .fail(function (jqXHR) {
            if (jqXHR.status === 200) {
                successHandler();
                return;
            }
            alert("Usuário ou senha inválidos!");
        });
}

function successHandler() {
    window.location = "/home";
}

function getFormFieldsValues() {
    const email = $("#email").val();
    const password = $("#password").val();

    return { email, password };
}
