$("#create-publication-form").on("submit", CreatePublication);

function CreatePublication(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    $.ajax({
        url: "/api/publication",
        method: "POST",
        data,
    })
        .done(successHandler)
        .fail(function (jqXHR) {
            if (jqXHR.status === 201) {
                successHandler();
                return;
            }

            Swal.fire("Ops...", "Falha ao criar publicação!", "error");
        });
}

function getFormFieldsValues() {
    const title = $("#title").val();
    const content = $("#content").val();

    return { title, content };
}

function successHandler() {
    Swal.fire("Sucesso!", "Publição criada com sucesso!", "success").then(() => {
        window.location = "/home";
    });
}
