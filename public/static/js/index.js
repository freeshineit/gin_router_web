(function () {
  $(".urlencoded").on("click", function () {
    var data = getFormValue();
    axios({
      method: "post",
      url: "/api/urlencoded_post?limit=10",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      },
      data: $.param(data)
    }).then((res) => {
      console.log(res.data);
      $(".urlencoded-msg").text(`success  ${new Date()}`);
    });
  });

  $(".json").on("click", function () {
    var data = getFormValue();
    axios({
      method: "post",
      url: "/api/json_post",
      headers: {
        "Content-Type": "application/json"
      },
      data
    }).then((res) => {
      console.log(res.data);
      $(".json-msg").text(`success  ${new Date()}`);
    });
  });

  $(".jsonandform").on("click", function () {
    var data = getFormValue();

    axios({
      method: "post",
      url: "/api/json_and_form_post",
      headers: {
        "Content-Type": "application/json"
      },
      data
    }).then((res) => {
      console.log(res.data);
      $(".jsonandform-msg").text(
        `success application/json data,  ${new Date()}`
      );
    });
  });

  $(".jsonandform2").on("click", function () {
    var data = getFormValue();

    axios({
      method: "post",
      url: "/api/json_and_form_post",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      },
      data: $.param(data)
    }).then((res) => {
      console.log(res.data);
      $(".jsonandform-msg").text(
        `success application/x-www-form-urlencoded data${new Date()}`
      );
    });
  });

  $(".xml_post").on("click", function () {
    var data = getFormValue();

    axios({
      method: "post",
      url: "/api/xml_post",
      headers: {
        "Content-Type": "application/xml"
      },
      data: `<xml><name>${data.name}</name><message>${data.message}</message><nick>${data.nick}</nick></xml>`
    }).then((res) => {
      $(".xml-msg").text(`success xml data ${new Date()}`);
    });
  });

  $(".file_upload").on("click", function () {
    // 单个文件上传
    // var fd = new FormData()
    // var file = document.getElementById('file')
    // fd.append('file', file.files[0])
    axios({
      method: "post",
      url: "/api/file_upload",
      headers: {
        "Content-Type": "application/form-data"
      },
      data: new FormData($("#multipleForm")[0])
    }).then((res) => {
      console.log(res.data);
      const urls = res.data.data.url.split(";");
      let imgHtml = "";

      for (let i = 0; i < urls.length; i++) {
        imgHtml += `<img style="width: 200px" src="/${urls[i]}" />`;
      }

      $(".file_upload-msg").html(
        `<div>success ${new Date()} 文件地址/${
          res.data.url
        } <div>${imgHtml}</div></div>`
      );
    });
  });

  // 成功
  $(".query").on("click", function () {
    axios
      .get("/api/query", {
        params: {
          message: "message",
          nick: "ShineShao"
        }
      })
      .then((res) => {
        console.log(res.data);
        $(".query-msg").text(`success ${new Date()}`);
      });
  });

  function getFormValue() {
    var data = {};
    var inputs = $("#form input");
    for (let i = 0; i < inputs.length; i++) {
      data[$(inputs[i]).attr("name")] = $(inputs[i]).val();
    }
    return data;
  }
})();
