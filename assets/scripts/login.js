$("#login-form").on("submit", login);

function login(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    console.log(data);

    $.ajax({
        url: "api/login",
        method: "POST",
        data,
    })
        .done(function () {
            window.location = "/home";
        })
        .fail(function () {
            alert("Usuário ou senha inválidos!");
        });
}

function getFormFieldsValues() {
    const email = $("#email").val();
    const password = $("#password").val();

    return { email, password };
}
