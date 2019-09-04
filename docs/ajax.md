---
layout: default
title: AJAX Compatability
description: Adding AJAX Compatability
has_children: true
nav_order: 25
---

# AJAX Compatability

Kosher has the ability to wait on AJAX requests just as it waits on normal HTTP requests.
{: .fs-6 .fw-300 }

This feature is a necessity for testing single-page applications
and applications that load data on-demand or heavily use AJAX in other ways. 

To take advantage of this kosher feature, you must modify your application to notify kosher of pending or active AJAX requests. Your app will provide a global
function that returns `true` if AJAX requests are pending/active and `false`, otherwise. This function must be added to the global `window` object and must be
named `ajaxPending`.

An example of such a function is shown below. This example adds an 'interceptor' to the [vue-resource API](https://github.com/pagekit/vue-resource) that maintains 
a list of unique IDs representing any AJAX requests. As requests are made, an ID is created and pushed to the `pendingAjaxRequests` array. When requests complete,
the ID is removed. Finally, the necessary `ajaxPending` function returns true if there are any IDs in the array.

```javascript

    window.pendingAjaxRequests = [];
    window.ajaxPending = () => window.pendingAjaxRequests.length > 0;

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