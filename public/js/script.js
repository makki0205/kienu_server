(function($){

  // base URL
  var BASE_URL = "http://" + window.location.host
  var UPLOAD_ENDPOINT = "/api/upload"

  // クリップボードへコピー
  var cb = new Clipboard(".copy")
  cb.on("success",function(e){
    alert("コピーしました！")
    e.clearSelection()
  })

  $("#submitButton").on("click",function(e){
    var form = $("#form").get()[0];

    var formData = new FormData(form);

    $.ajax({
      async: true,
      url: BASE_URL + UPLOAD_ENDPOINT,
      type: "POST",
      dataType: "json",
      data: formData,
      processData: false,
      contentType: false,
      crossDomain: true,
      cache: false,
      xhr: function(){
        var XHR = $.ajaxSettings.xhr();
        if(XHR.upload){
          XHR.upload.addEventListener("progress",function(e){
            var progress = parseInt(e.loaded/e.total*100);
            $(".progress-bar").css({ width: progress + "%" })
          },false)
        }
        return XHR;
      }
    })
    .done(function(res){
      var downloadUrl = BASE_URL + res.download_url
      $("#download-url").val(downloadUrl)
      $(".downloadedElement").show()
      $(".progress-bar").hide().css({ width: "0%" })
        console.log(res.download_url)
    })
    .fail(function(jqXHR,textStatus,error){
      console.log("ERROR: ",jqXHR,textStatus,error)
    });

    return false

  })

})(jQuery)