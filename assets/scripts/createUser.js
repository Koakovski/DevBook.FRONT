$("#create-user-form").on("submit", createUser);

function createUser(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    if (data.password != data.confirmPassword) {
        alert("Senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "api/user",
        method: "POST",
        data,
    });
}

function getFormFieldsValues() {
    const name = $("#name").val();
    const email = $("#email").val();
    const nickname = $("#nickName").val();
    const password = $("#password").val();
    const confirmPassword = $("#confirmPassword").val();

    return { name, email, nickname, password, confirmPassword };
}
