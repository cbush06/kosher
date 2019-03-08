package reporttemplates

// GetSimpleTemplate returns a simplified template for Kosher reports that is tailored
// to be imported into Word and still appear correctly
func GetSimpleTemplate() string {
	return `<html lang="en">
    <head>
        <meta charset="utf-8" />
        <title></title>
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous" />
		<style type="text/css">
			.bg-success { background-color: #c7fac9 !important; }
			.bg-danger { background-color: #f9c8c9 !important; }
			.bg-info { background-color: #c8f9fb !important; }
            .table td, .table th { border: none !important; }
		</style>
    </head>
    <body>
        <div class="jumbotron jumbotron-fluid bg-secondary text-white pt-3 pb-3">
            <div class="container-fluid">
                <div class="row">
                    <div class="col-sm">
                        <h1 class="display-4">{{.ProjectName}}</h1>
                    </div>
                    <div class="col-sm ml-auto text-right">
                        <h2 class="mb-0">
                            <span class="badge badge-success">Passed: {{.StepsPassed}}</span>
                            <span class="badge badge-danger">Failed: {{.StepsFailed}}</span>
                            <span class="badge badge-info">Pending: {{.StepsPending}}</span>
                            <span class="badge badge-warning  text-white">Skipped: {{.StepsSkipped}}</span>
                        </h2>
                    </div>
                </div>
                <div class="row">
                    <div class="col-sm">
                        <h3>Kosher Test Results</h2>
                    </div>
                    <div class="col-sm ml-auto text-right">
                        <h3>{{.Timestamp}}</h2>
                    </div>
                </div>
            </div>
        </div>

        <div class="container-fluid">
            <table class="table border bg-light">
                <thead>
                    <tr>
                        <th colspan="4" class="table-secondary">Metadata</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <th>App Version:</th>
                        <td>{{.AppVersion}}</td>
                        <th>Environment:</th>
                        <td>{{.Environment}}</td>
                    </tr>
                    <tr>
						<th>Platform:</th>
						<td>{{.Platform}}</td>
                        <th>Browser:</th>
                        <td>{{.Browser}}</td>
					</tr>
					<tr>
						<th>OS:</th>
						<td>{{.OS}}</td>
						<th>Total Run Time:</th>
						<td>{{.RunTime}}</td>
                </tbody>
            </table>
		</div>
		
		<p>&nbsp;</p>

        <div class="container-fluid">
            {{range .Features}}
            
                <!-- FEATURE BLOCK -->
                <h1 class="{{if (gt .ElementsFailed 0)}}bg-danger{{else if (gt .ElementsPending 0)}}bg-inf{{else}}bg-success{{end}}">
                    <strong>{{if (gt .ElementsFailed 0)}}(FAILED){{else if (gt .ElementsPending 0)}}(PENDING){{else}}(PASSED){{end}} Feature:</strong> {{.Name}}
                </h1>
				<p class="ml-4 mr-4" style="font-size: 1rem;"><pre>{{.GetTrimmedDescription}}</pre></p>
				<p>&nbsp;</p>
                {{range .Elements}}

                    <!-- SCENARIO, SCENARIO OUTLINE, OR BACKGROUND BLOCK -->
                    <div class="ml-4 mr-4">
						<table class="table table-sm table-borderless">
							<thead>
								<th colspan="2" class="{{if (gt .StepsFailed 0)}}bg-danger{{else if (gt .StepsSkipped 0)}}bg-info{{else}}bg-success{{end}}">
									<h2><strong>{{if (gt .StepsFailed 0)}}(FAILED){{else if (gt .StepsSkipped 0)}}(PENDING){{else}}(PASSED){{end}} {{.Keyword}}:</strong> {{.Name}}</h2>
								</th>
							</thead>
							<tbody>
								<tr><td colspan="2" class="pl-4 pr-4"><pre>{{.GetTrimmedDescription}}</pre></td></tr>

                                <!-- STEP BLOCK -->
                                {{range .Steps}}
                                <tr>
                                    <td class="{{if (eq .Result.Status "passed")}}text-success{{else if (eq .Result.Status "failed")}}text-danger{{else if (eq .Result.Status "skipped")}}text-warn{{else if (eq .Result.Status "undefined")}}text-info{{end}}">
										<strong>{{if (eq .Result.Status "passed")}}(PASSED){{else if (eq .Result.Status "failed")}}(FAILED){{else if (eq .Result.Status "skipped")}}(SKIPPED){{else if (eq .Result.Status "undefined")}}(PENDING){{end}} {{.Keyword}}</strong>{{.Name}}</span>
									</td>
                                    </td>
                                    <td style="width: 10%;">
                                        {{.Result.GetDurationInSeconds}}
                                    </td>
                                </tr>
                                {{if (eq .Result.Status "failed")}}
                                    <tr>
                                        <td colspan="2" class="bg-danger">
                                            {{.Result.Error}}
                                        </td>
                                    </tr>
                                {{else if (eq .Result.Status "undefined")}}
                                    <tr>
                                        <td colspan="2" class="bg-info">
                                            Could not match step at <code>{{.Match.Location}}</code>
                                        </td>
                                    </tr>
                                {{end}}
                                {{end}}
                            </tbody>
                        </table>
                    </div>
                {{end}}
			{{end}}
        </div>
    </body>
</html>`
}