---
layout: default
title: AJAX Compatability
description: Adding AJAX Compatability
has_children: true
nav_order: 25
---

# AJAX Compatability

In order for kosher to work properly with AJAX, your app should provide a function that returns true if AJAX requests are pending/active.

```javascript

    window.pendingAjaxRequests = [];
    window.ajaxPending = () => window.pendingAjaxRequests.length == 0;

    Vue.http.interceptor.before = function(request) {
        const requestUkId = request.getUrl() + "#" + Date.now();
        window.pendingAjaxRequests.push(requestUkId);

        // e is a ProgressEvent -- https://developer.mozilla.org/en-US/docs/Web/API/ProgressEvent
        request.downloadProgress = function(e) {
            console.log(e)
            if(e.loaded == e.total) {
                _.remove(window.pendingAjaxRequests, requestUkId);
            }
        };
    };

```