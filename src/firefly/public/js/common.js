var AppendUrlToShareContent, ConfirmDelete, FindElementInArray, FormatNowTime, GenerateShareContent, IsUsernameMentioned, RemoveElementInArray, ReplaceBookTitleToString, ReplaceLinkString, ReplaceMentionToLink, ReplaceReferToLink, RequestAjax, RequestAjaxWithParam, ScrollToElement, getAvatarUrl, getQueryString;

window.scoreHints = ['难看', '不好看', '一般', '好书', '神作'];

window.scoreNoRatedMsg = "尚无评分";

getQueryString = function(name) {
  var r, reg;
  reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
  r = window.location.search.substr(1).match(reg);
  if (r !== null) {
    return escape(r[2]);
  }
  return null;
};

getAvatarUrl = function(url, size) {
  return url.replace("%d", size);
};

RequestAjax = function(type, url, data, successCallback, failCallback, timeoutCallback, beforeAction, afterAction, dontAlertOnStatusCode, async) {
  var rawData;
  data.csrf_token = window.csrfToken;
  rawData = {
    Path: url,
    Data: data
  };
  $.ajax({
    async: async != null ? async : true,
    type: type,
    url: encodeURI(url),
    cache: false,
    data: data,
    statusCode: {
      400: (function() {
        if (!dontAlertOnStatusCode) {
          return alert("400: 请求不正确");
        }
      }),
      404: (function() {
        if (!dontAlertOnStatusCode) {
          return alert("404: 该资源不存在");
        }
      }),
      500: (function() {
        if (!dontAlertOnStatusCode) {
          return alert("500: 服务器遇到一个内部错误，请稍等一会再试试");
        }
      })
    },
    success: (function(data) {
      var _ref;
      if (data.Result === true) {
        if (successCallback != null) {
          successCallback(data, rawData);
        } else {
          location.reload();
        }
      } else {
        if (failCallback != null) {
          failCallback(data, rawData);
        } else {
          alert((_ref = data.ErrorMsg) != null ? _ref : "网络发生故障，请稍后重新尝试");
        }
      }
    }),
    dataType: "json",
    timeout: 4000,
    error: (function(jqXHR, textStatus, errorThrown) {
      if (textStatus === "timeout") {
        if (timeoutCallback != null) {
          timeoutCallback(errorThrown);
        } else {
          if (typeof failCallback === "function") {
            failCallback(data, rawData);
          }
        }
      }
      if (textStatus === "error") {
        if (typeof failCallback === "function") {
          failCallback(errorThrown);
        }
      }
    }),
    beforeSend: (function(jqXHR, settings) {
      if (typeof beforeAction === "function") {
        beforeAction(jqXHR, settings);
      }
    }),
    complete: (function(jqXHR, textStatus) {
      if (typeof afterAction === "function") {
        afterAction(jqXHR, textStatus);
      }
    })
  });
};

RequestAjaxWithParam = function(options) {
  var rawData, _ref, _ref1, _ref2, _ref3;
  if (options.data != null) {
    options.data.csrf_token = window.csrfToken;
  } else {
    options.data = {
      csrf_token: window.csrfToken
    };
  }
  rawData = {
    Path: (_ref = options.url) != null ? _ref : "/",
    Data: options.data
  };
  $.ajax({
    async: (_ref1 = options.async) != null ? _ref1 : true,
    type: (_ref2 = options.type) != null ? _ref2 : "GET",
    url: encodeURI((_ref3 = options.url) != null ? _ref3 : "/"),
    cache: false,
    data: options.data,
    statusCode: {
      400: (function() {
        if (!options.dontAlertOnStatusCode) {
          return alert("400: 请求不正确");
        }
      }),
      404: (function() {
        if (!options.dontAlertOnStatusCode) {
          return alert("404: 该资源不存在");
        }
      }),
      500: (function() {
        if (!options.dontAlertOnStatusCode) {
          return alert("500: 服务器遇到一个内部错误，请稍等一会再试试");
        }
      })
    },
    success: (function(data) {
      var _ref4;
      if (data.Result === true) {
        if (options.successCallback != null) {
          options.successCallback(data, rawData);
        } else {
          location.reload();
        }
      } else {
        if (options.failCallback != null) {
          options.failCallback(data, rawData);
        } else {
          alert((_ref4 = data.ErrorMsg) != null ? _ref4 : "网络发生故障，请稍后重新尝试");
        }
      }
    }),
    dataType: "json",
    timeout: 4000,
    error: (function(jqXHR, textStatus, errorThrown) {
      if (textStatus === "timeout") {
        if (options.timeoutCallback != null) {
          options.timeoutCallback(errorThrown);
        } else {
          if (typeof options.failCallback === "function") {
            options.failCallback(data, rawData);
          }
        }
      }
      if (textStatus === "error") {
        if (typeof options.failCallback === "function") {
          options.failCallback(errorThrown);
        }
      }
    }),
    beforeSend: (function(jqXHR, settings) {
      if (typeof options.beforeAction === "function") {
        options.beforeAction(jqXHR, settings);
      }
    }),
    complete: (function(jqXHR, textStatus) {
      if (typeof options.afterAction === "function") {
        options.afterAction(jqXHR, textStatus);
      }
    })
  });
};

ConfirmDelete = function(sender, okAction, cancelAction, actionString) {
  if ((actionString != null ? actionString.length : void 0) > 0) {
    $("h4.label_Title").text("确认" + actionString);
    $(".modal-body").children("div").text("确定要" + actionString + "吗？");
    $(".btn_confirm").text(actionString);
  }
  $("#modal_confirm").modal();
  $(".btn_cancel").unbind("click").click(function(event) {
    if (typeof cancelAction === "function") {
      cancelAction(sender);
    }
  });
  $(".btn_confirm").unbind("click").click(function(event) {
    if (typeof okAction === "function") {
      okAction(sender);
    }
  });
};

FindElementInArray = function(array, element) {
  var e, i, returnValue, _i, _len;
  returnValue = -1;
  for (i = _i = 0, _len = array.length; _i < _len; i = ++_i) {
    e = array[i];
    if (e === element) {
      returnValue = i;
    }
  }
  return returnValue;
};

RemoveElementInArray = function(array, element) {
  var i, ser, _i, _len;
  for (i = _i = 0, _len = array.length; _i < _len; i = ++_i) {
    ser = array[i];
    if (ser === element) {
      array.splice(i, 1);
    }
  }
  return array;
};

GenerateShareContent = function(options) {
  var content, i, shareStringPrefix, star, _i, _ref;
  content = "";
  star = "";
  for (i = _i = 1, _ref = options.score; 1 <= _ref ? _i <= _ref : _i >= _ref; i = 1 <= _ref ? ++_i : --_i) {
    star += "★";
  }
  switch (options.type) {
    case "cmt":
      shareStringPrefix = "" + options.user + "给《" + options.bookTitle + "》" + star + "：";
      if (options.user !== "我") {
        shareStringPrefix = "喜欢" + options.user + "为《" + options.bookTitle + "》写的书评：";
      }
      content = shareStringPrefix + options.content.substr(0, 120 - shareStringPrefix.length);
      break;
    case "review":
      shareStringPrefix = "" + options.user + "读过《" + options.bookTitle + "》给" + star + "，并写了书评“";
      if (options.user !== "我") {
        shareStringPrefix = "喜欢" + options.user + "为《" + options.bookTitle + "》写的书评“";
      }
      content = shareStringPrefix + options.reviewTitle.substr(0, 120 - shareStringPrefix.length - 9) + "”";
  }
  return content;
};

AppendUrlToShareContent = function(options) {
  var fullContent;
  fullContent = "";
  switch (options.type) {
    case "cmt":
      fullContent = options.content + location.href;
      break;
    case "review":
      fullContent = options.content + ("查看全文( http://www.shanpow.com" + options.relatedUrl + " )");
  }
  return fullContent;
};

$.fn.MoveToEnd = function() {
  var len, obj, sel;
  obj = $(this)[0];
  obj.focus();
  len = 0;
  if (obj.value != null) {
    len = obj.value.length;
  } else {
    len = obj.textContent.length;
  }
  if (document.selection) {
    sel = obj.createTextRange();
    sel.moveStart('character', len);
    sel.collapse();
    sel.select();
  } else if (typeof obj.selectionStart === 'number' && typeof obj.selectionEnd === 'number') {
    obj.selectionStart = obj.selectionEnd = len;
  }
};

$.fn.InsertTextOnCursor = function(str) {
  var cursorPos, endPos, obj, sel, startPos, tmpStr;
  obj = $(this)[0];
  if (document.selection) {
    sel = document.selection.createRange();
    sel.text = str;
  } else if (typeof obj.selectionStart === 'number' && typeof obj.selectionEnd === 'number') {
    startPos = obj.selectionStart;
    endPos = obj.selectionEnd;
    cursorPos = startPos;
    tmpStr = obj.value;
    obj.value = tmpStr.substring(0, startPos) + str + tmpStr.substring(endPos, tmpStr.length);
    cursorPos += str.length;
    obj.selectionStart = obj.selectionEnd = cursorPos;
  } else {
    obj.value += str;
  }
};

String.prototype.InsertStringAtIndex = function(str, index) {
  var array, i, s, _i, _len;
  array = this.split("");
  for (i = _i = 0, _len = array.length; _i < _len; i = ++_i) {
    s = array[i];
    if (i === index) {
      array[i] = "" + str + s;
    }
  }
  return array.join("");
};

String.prototype.Trim = function() {
  return this.replace(/(^\s+)|(\s+$)/g, "");
};

String.prototype.RemoveHtmlTag = function() {
  return this.replace(/<[^<>]+?>/g, "");
};

String.prototype.RemoveScriptTag = function() {
  var str;
  str = this.replace(/<script>/g, "");
  return str.replace(/<\/script>/g, "");
};

String.prototype.Escape2Html = function() {
  var arrEntities;
  arrEntities = {
    'lt': '<',
    'gt': '>',
    'nbsp': ' ',
    'amp': '&',
    'quot': '"'
  };
  return this.replace(/&(lt|gt|nbsp|amp|quot);/ig, (function(all, t) {
    return arrEntities[t];
  }));
};

Array.prototype.RemoveDuplicateElement = function() {
  var e, k, keyArray, newArray, oldArray, _i, _j, _len, _len1;
  oldArray = this || [];
  newArray = {};
  keyArray = [];
  for (_i = 0, _len = oldArray.length; _i < _len; _i++) {
    e = oldArray[_i];
    if (typeof newArray[e] === "undefined") {
      newArray[e] = 1;
      keyArray.push(e);
    }
  }
  oldArray.length = 0;
  for (_j = 0, _len1 = keyArray.length; _j < _len1; _j++) {
    k = keyArray[_j];
    oldArray[oldArray.length] = k;
  }
  return oldArray;
};

ReplaceReferToLink = function(str, refers) {
  var divToInsert, indexStr, preA, refer, referContent, referindexes, _i, _j, _len, _len1;
  if (str.indexOf("<p>") !== 0) {
    str = "<p>" + str + "</p>";
  }
  referindexes = str.match(/#(\d+)/g);
  if (referindexes != null) {
    for (_i = 0, _len = referindexes.length; _i < _len; _i++) {
      indexStr = referindexes[_i];
      preA = "</p><a href='#' class='referindex' data-index='" + (indexStr.replace('#', '')) + "'>";
      divToInsert = "";
      for (_j = 0, _len1 = refers.length; _j < _len1; _j++) {
        refer = refers[_j];
        if (parseInt(refer.Index) === parseInt(indexStr.replace("#", ""))) {
          referContent = "" + refer.Author.Nickname + "：" + refer.Content;
          if (refer.Content.length > 200) {
            referContent = refer.Content.substr(0, 200) + "...";
          }
          divToInsert += "<div class='refercontent' data-refercontent='" + referContent + "'></div>";
        }
      }
      str = str.replace("" + indexStr, "" + preA + indexStr + "</a>" + divToInsert + "<p>");
    }
  }
  return str;
};

ReplaceMentionToLink = function(str) {
  var mentionStr, mentionindexes, nickname, _i, _len;
  if (str.indexOf("<p>") !== 0) {
    str = "<p>" + str + "</p>";
  }
  mentionindexes = str.match(/@(.*?)[\s|\n|\r]/g);
  if (mentionindexes != null) {
    for (_i = 0, _len = mentionindexes.length; _i < _len; _i++) {
      mentionStr = mentionindexes[_i];
      nickname = mentionStr.replace("@", "");
      str = str.replace("@" + nickname, "</p><p>@</p><a href='/people/" + nickname + "' class='mentionindex' target='blank'>" + nickname + "</a><p>");
    }
  }
  return str;
};

ReplaceLinkString = function(str) {
  var link, links, reg, _i, _len;
  str = "<p>" + str + "</p>";
  links = str.match(/(http:\/\/\S*)\s/g);
  links = links != null ? links.RemoveDuplicateElement() : void 0;
  if (links != null) {
    for (_i = 0, _len = links.length; _i < _len; _i++) {
      link = links[_i];
      reg = new RegExp(link, "gi");
      str = str.replace(reg, "</p><a href='" + (link.Trim()) + "' target='blank'>" + (link.Trim()) + "</a>" + (link.charAt(link.length - 1)) + "<p>");
    }
  }
  return str;
};

ReplaceBookTitleToString = function(str) {
  var bookname, booktitle, booktitles, _i, _len;
  if (str.indexOf("<p>") !== 0) {
    str = "<p>" + str + "</p>";
  }
  booktitles = str.match(/《(.*?)》/g);
  if (booktitles != null) {
    for (_i = 0, _len = booktitles.length; _i < _len; _i++) {
      booktitle = booktitles[_i];
      bookname = booktitle.replace("《", "").replace("》", "");
      str = str.replace("《" + bookname + "》", "</p>《<a href='/search?q=" + bookname + "' class='booktitle' target='blank'>" + bookname + "</a>》<p>");
    }
  }
  return str;
};

$.fn.InsertLinkIntoElement = function(refers) {
  var element, str, textStr;
  element = $(this);
  textStr = element.text().replace(/</g, "&lt;").replace(/>/g, "&gt;");
  str = ReplaceLinkString(textStr);
  str = ReplaceMentionToLink(str);
  str = ReplaceBookTitleToString(str);
  str = ReplaceReferToLink(str, refers);
  element.html(str);
  element.children("p").each(function(index) {
    $(this).text($(this).html().Escape2Html());
    $(this).contents().unwrap();
  });
  if (str === textStr) {
    element.text(element.html());
  }
  element.find(".refercontent").each(function(index) {
    $(this).text($(this).data("refercontent"));
  });
};

$.fn.TextWithoutReferContent = function() {
  var element, str;
  element = $(this);
  str = element.text();
  element.find(".refercontent").each(function(index) {
    str = str.replace($(this).text(), "");
  });
  return str;
};

IsUsernameMentioned = function(content, username) {
  var str;
  str = "@" + username;
  if (content == null) {
    return -1;
  }
  return content.indexOf(str);
};

FormatNowTime = function() {
  var now;
  now = new Date();
  return "" + (now.getFullYear()) + "-" + (now.getMonth() + 1) + "-" + (now.getDate()) + "&nbsp;" + (now.getHours()) + ":" + (now.getMinutes()) + ":" + (now.getSeconds());
};

ScrollToElement = function(element) {
  var y;
  y = element.offset().top;
  $(document).scrollTop(y - 50);
};

$(document).ready(function(event) {
  var sys, ua;
  sys = {};
  ua = navigator.userAgent.toLowerCase();
  if (window.ActiveXObject) {
    sys.ie = ua.match(/msie ([\d.]+)/)[1];
  }
  if (parseFloat(sys.ie) < 8.0) {
    location.href = "/upgradebrowser";
  }
  $(window).scroll(function(event) {
    if ($(window).scrollTop() > 0) {
      $(".backToTop").fadeIn();
    } else {
      $(".backToTop").fadeOut();
    }
  });
  $(".backToTop").unbind("click").click(function(event) {
    $(window).scrollTop(0);
  });
});
