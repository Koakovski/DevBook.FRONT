$("#create-user-form").on("submit", createUser);

function createUser(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    if (data.password != data.confirmPassword) {
        alert("Senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "/api/user",
        method: "POST",
        data,
    })
        .done(successHandler)
        .fail(function (jqXHR) {
            if (jqXHR.status === 201) {
                successHandler();
                return;
            }

            alert("Falha ao cadastror o usuário!");
        });
}

function successHandler() {
    alert("Usuário cadastrado com sucesso!");
    window.location = "/login";
}

function getFormFieldsValues() {
    const name = $("#name").val();
    const email = $("#email").val();
    const nickName = $("#nickName").val();
    const password = $("#password").val();
    const confirmPassword = $("#confirmPassword").val();

    return { name, email, nickName, password, confirmPassword };
}
