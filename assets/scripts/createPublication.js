$("#create-publication-form").on("submit", CreatePublication);

function CreatePublication(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    $.ajax({
        url: "api/publication",
        method: "POST",
        data,
    })
        .done(successHandler)
        .fail(function (jqXHR) {
            if (jqXHR.status === 201) {
                successHandler();
                return;
            }

            alert("Falha ao criar publicação!");
        });
}

function getFormFieldsValues() {
    const title = $("#title").val();
    const content = $("#content").val();

    return { title, content };
}

function successHandler() {
    window.location = "/home";
}