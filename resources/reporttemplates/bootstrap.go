package reporttemplates

// GetBootstrapTemplate returns the Bootstrap template for Kosher reports.
func GetBootstrapTemplate() string {
	return `<html lang="en">
    <head>
        <meta charset="utf-8" />
        <title></title>
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous" />
		<link href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet" integrity="sha384-wvfXpqpZZVQGK6TAh5PVlGOfQNHSoD2xbE+QkPxCAFlNEevoEH3Sl0sibVcOQVnN" crossorigin="anonymous">
		<style type="text/css">
			.bg-success { background-color: #c7fac9 !important; }
			.bg-danger { background-color: #f9c8c9 !important; }
			.bg-info { background-color: #c8f9fb !important; }
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
            <div class="row">
                <div class="col-6">
                    <div class="card bg-light mb-3">
                        <div class="card-body">
                            <h4 class="card-title">Scenarios ({{.TotalElements}})</h4>
                            <div class="ml-auto mr-auto" style="width: 400px; height: 400px;">
                                <canvas id="scenariosChart" width="400" height="400"></canvas>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-6">
                    <div class="card bg-light mb-3">
                        <div class="card-body">
                            <h4 class="card-title">Steps ({{.TotalSteps}})</h4>
                            <div class="ml-auto mr-auto" style="width: 400px; height: 400px;">
                                <canvas id="stepsChart" width="400" height="400"></canvas>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col">
                    <div class="card bg-light mb-3">
                        <div class="card-header">
                            Metadata
                        </div>
                        <div class="card-body">
                            <div class="container-fluid">
                                <div class="row">
                                    <span class="col-2"><strong>App Version:</strong></span>
                                    <span class="col-4">{{.AppVersion}}</span>

                                    <span class="col-2"><strong>Environment:</strong></span>
                                    <span class="col-4">{{.Environment}}</span>

									<span class="col-2"><strong>Platform:</strong></span>
									<span class="col-4">{{.Platform}}</span>

                                    <span class="col-2"><strong>Browser:</strong></span>
									<span class="col-4">{{.Browser}}</span>
									
									<span class="col-2"><strong>OS:</strong></span>
									<span class="col-4">{{.OS}}</span>

									<span class="col-2"><strong>Total Run Time:</strong></span>
									<span class="col-4">{{.RunTime}}</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
			{{range .Features}}

			<!-- FEATURE BLOCK -->
			<div class="row">
                <div class="col">
                    <div class="card mb-3">
						<div class="card-header {{if (gt .ElementsFailed 0)}}bg-danger{{else if (gt .ElementsPending 0)}}bg-info{{else}}bg-success{{end}}">
							{{ if (gt (len .Tags) 0) }}
							<small>
								{{range .Tags}}
									{{.Name}}&nbsp;
								{{end}}
							</small>
							<br />
							{{end}}

							<strong>Feature:</strong> {{.Name}}
						</div>
						<div class="card-body">
							<div class="card-title container-fluid">
								<div class="row"><div class="col"><pre>{{.GetTrimmedDescription}}</pre></div></div>
							</div>
							<div class="container-fluid">
								{{range .Elements}}

								<!-- SCENARIO, SCENARIO OUTLINE, OR BACKGROUND BLOCK -->
								<div class="row">
									<div class="col">
										<div class="card mb-3">
											<div class="card-header {{if (gt .StepsFailed 0)}}bg-danger{{else if (gt .StepsSkipped 0)}}bg-info{{else}}bg-success{{end}}">
												{{ if (gt (len .Tags) 0) }}
												<small>
													{{range .Tags}}
														{{.Name}}&nbsp;
													{{end}}
												</small>
												<br />
												{{end}}
												
												<strong>{{.Keyword}}:</strong> {{.Name}}
											</div>
											<div class="card-body">
												<div class="card-title container-fluid">
													<div class="row"><div class="col"><pre>{{.GetTrimmedDescription}}</pre></div></div>
												</div>
												<div class="container-fluid">

													<!-- STEP BLOCK -->
													{{range .Steps}}
														<div class="row">
															<div class="col-11">
																{{if (eq .Result.Status "passed")}}
																	<span class="text-success"><i class="fa fa-check-square"></i></span>
																{{else if (eq .Result.Status "failed")}}
																	<span class="text-danger"><i class="fa fa-times-circle"></i></span>
																{{else if (eq .Result.Status "skipped")}}
																	<span class="text-warning"><i class="fa fa-exclamation-triangle"></i></span>
																{{else if (eq .Result.Status "undefined")}}
																	<span class="text-info"><i class="fa fa-question-circle"></i></span>
																{{end}}

																<strong>{{.Keyword}}</strong>{{.Name}}
															</div>
															<div class="col-1">{{.Result.GetDurationInSeconds}}</div>
														</div>

														{{if (eq .Result.Status "failed")}}
															<div class="row">
																<div class="col alert alert-danger">
																	{{.Result.Error}}
																</div>
															</div>
														{{else if (eq .Result.Status "undefined")}}
															<div class="row">
																<div class="col alert alert-info">
																	Could not match step at <code>{{.Match.Location}}</code>
																</div>
															</div>
														{{end}}
													{{end}}
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
							</div>
						</div>
					</div>
				</div>
			</div>
			{{end}}
        </div>


        <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js" integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8=" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.24.0/moment.min.js" integrity="sha256-4iQZ6BVL4qNKlQ27TExEhBN1HFPvAvAMbFavKKosSWQ=" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.7.3/Chart.min.js" integrity="sha256-oSgtFCCmHWRPQ/JmR4OoZ3Xke1Pw4v50uh6pLcu+fIc=" crossorigin="anonymous"></script>
        <script type="text/javascript">
            var ctx = document.getElementById("scenariosChart");
            new Chart(ctx, {
                type: 'doughnut',
                data: {
                    labels: ['Passed', 'Failed'],
                    datasets: [{
                        data: [{{.ElementsPassed}}, {{.ElementsFailed}}],
                        backgroundColor: [
                            'rgba(0, 255, 0, 0.2)',
                            'rgba(255, 0, 0, 0.2)'
                        ],
                        borderColor: [
                            'rgba(0, 255, 0, 0.2)',
                            'rgba(255, 0, 0, 0.2)'
                        ],
                        borderWidth: 1
                    }]
				},
				options: {
					legend: {
						labels: {
							generateLabels: labelFunc
						}
					}
				}
            });

            var ctx = document.getElementById("stepsChart");
            new Chart(ctx, {
                type: 'doughnut',
                data: {
                    labels: ['Passed', 'Failed', 'Pending', 'Skipped'],
                    datasets: [{
                        data: [{{.StepsPassed}}, {{.StepsFailed}}, {{.StepsPending}}, {{.StepsSkipped}}],
                        backgroundColor: [
                            'rgba(0, 255, 0, 0.2)',
                            'rgba(255, 0, 0, 0.2)',
                            'rgba(0, 250, 255, 0.2)',
                            'rgba(255, 200, 0, 0.2)'
                        ],
                        borderColor: [
                            'rgba(0, 255, 0, 1)',
                            'rgba(255, 0, 0, 1)',
                            'rgba(0, 250, 255, 1)',
                            'rgba(255, 200, 0, 1)'
                        ],
                        borderWidth: 1
                    }]
				},
				options: {
					legend: {
						labels: {
							generateLabels: labelFunc
						}
					}
				}
			});
			
			function labelFunc(chart) {
				var data = chart.data;
				if (data.labels.length && data.datasets.length) {
					return data.labels.map(function(label, i) {
						var meta = chart.getDatasetMeta(0);
						var ds = data.datasets[0];
						var arc = meta.data[i];
						var custom = arc && arc.custom || {};
						var getValueAtIndexOrDefault = Chart.helpers.getValueAtIndexOrDefault;
						var arcOpts = chart.options.elements.arc;
						var fill = custom.backgroundColor ? custom.backgroundColor : getValueAtIndexOrDefault(ds.backgroundColor, i, arcOpts.backgroundColor);
						var stroke = custom.borderColor ? custom.borderColor : getValueAtIndexOrDefault(ds.borderColor, i, arcOpts.borderColor);
						var bw = custom.borderWidth ? custom.borderWidth : getValueAtIndexOrDefault(ds.borderWidth, i, arcOpts.borderWidth);

						// We get the value of the current label
						var value = chart.config.data.datasets[arc._datasetIndex].data[arc._index];

						return {
							// Add value to label
							text: label + " (" + value + ")",
							fillStyle: fill,
							strokeStyle: stroke,
							lineWidth: bw,
							hidden: isNaN(ds.data[i]) || meta.data[i].hidden,
							index: i
						};
					});
				} else {
					return [];
				}
			}
        </script>
    </body>
</html>`
}
