package reporttemplates

// GetAxeTemplate returns the template used to generate Axe Accessibility Scan reports
func GetAxeTemplate() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="utf-8" />
		<title></title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
		<link href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet" integrity="sha384-wvfXpqpZZVQGK6TAh5PVlGOfQNHSoD2xbE+QkPxCAFlNEevoEH3Sl0sibVcOQVnN" crossorigin="anonymous">
		<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
	</head>
	<body>
		<div class="jumbotron jumbotron-fluid bg-secondary text-white pt-3 pb-3">
			<div class="container-fluid">
				<div class="row">
					<div class="col-sm">
						<h1 class="display-4">{{.ProjectName}}</h1>
					</div>
					<div class="col-sm ml-auto text-right">
						<h4 class="mb-0">
							<span class="badge badge-dark">Needs Review: {{.IncompleteCount}}</span>
							<span class="badge badge-info">Minor: {{.MinorCount}}</span>
							<span class="badge badge-primary">Moderate: {{.ModerateCount}}</span>
							<span class="badge badge-warning  text-white">Serious: {{.SeriousCount}}</span>
							<span class="badge badge-danger">Critical: {{.CriticalCount}}</span>
						</h4>
					</div>
				</div>
				<div class="row">
					<div class="col-sm">
						<h3>Kosher Accessibility Results</h2>
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
							<h4 class="card-title">Results by Outcome</h4>
							<div class="ml-auto mr-auto" style="width: 400px; height: 400px;">
								<canvas id="resultsByOutcomePie" width="400" height="400"></canvas>
							</div>
						</div>
					</div>
				</div>
				<div class="col-6">
					<div class="card bg-light mb-3">
						<div class="card-body">
							<h4 class="card-title">Violations by Severity</h4>
							<div class="ml-auto mr-auto" style="width: 400px; height: 400px;">
								<canvas id="violationsBySeverityBar" width="400" height="400"></canvas>
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
	
									<span class="col-2"><strong>Axe Version:</strong></span>
									<span class="col-4">{{.AxeVersion}}</span>
	
									<span class="col-2"><strong>OS:</strong></span>
									<span class="col-4">{{.OS}}</span>
	
									<span class="col-2"><strong>Rule Sets:</strong></span>
									<span class="col-4">{{.RuleSets}}</span>
	
									<span class="col-2"><strong>Environment:</strong></span>
									<span class="col-4">{{.Environment}}</span>
	
									<span class="col-2"><strong>Threshold:</strong></span>
									<span class="col-4">{{.ImpactThreshold}}</span>
	
									<span class="col-2"><strong>Browser:</strong></span>
									<span class="col-4">{{.Browser}}</span>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
	
			{{range .AxeScans}}
			{{if (or (gt (len .Violations) 0) (gt (len .Incomplete) 0))}}
			<!-- PAGE SCAN RESULTS -->
			<div class="row mb-3">
				<div class="col">
					<div class="card">
						<div class="card-header">
							<strong>{{.Title}}</strong><br/>
							<small class="text-muted">{{.URL}}</small>
						</div>
						<div class="card-body">
							<div class="container-fluid">
								<!-- VIOLATION RULES OF PAGE -->
								{{range .Violations}}
								<div class="row mb-3">
									<div class="col">
										<div class="card bg-light mb-3">
											<div class="card-header">
												<div class="container-fluid">
													<div class="row">
														<div class="col-sm-9">
															<strong>{{.Help}}</strong><br/>
															<small class="text-muted">{{.Description}}</small>
														</div>
														<div class="col-sm-3 text-right">
															{{if (eq .Impact "minor")}}
																<span class="badge badge-info">Minor</span>
															{{else if (eq .Impact "moderate")}}
																<span class="badge badge-primary">Moderate</span>
															{{else if (eq .Impact "serious")}}
																<span class="badge badge-warning text-white">Serious</span>
															{{else if (eq .Impact "critical")}}
																<span class="badge badge-danger">Critical</span>
															{{end}}<br />
															<a href="{{.HelpURL}}" target="_blank"><small><i class="fa fa-external-link"></i> Learn more</small></a>
														</div>
													</div>
												</div>
											</div>
											<div class="card-body">
												<div class="card-columns">
													{{range .Nodes}}
													<div class="card">
														<div class="card-body">
															<p><strong>{{(last .Target)}}</strong></p>
															<p class="bg-secondary"><code class="text-light">{{.HTML}}</code></p>
															{{.GetPrettyFailureSummary}}
														</div>
													</div>
													{{end}}
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
	
								<!-- VIOLATION RULES NEEDING REVIEW OF PAGE -->
								{{range .Incomplete}}
								<div class="row mb-3">
										<div class="col">
											<div class="card bg-light mb-3">
												<div class="card-header">
													<div class="container-fluid">
														<div class="row">
															<div class="col-sm-9">
																<strong>{{.Help}}</strong><br/>
																<small class="text-muted">{{.Description}}</small>
															</div>
															<div class="col-sm-3 text-right">
																<span class="badge badge-secondary">Needs Review</span><br/>
																<a href="{{.HelpURL}}" target="_blank"><small><i class="fa fa-external-link"></i> Learn more</small></a>
															</div>
														</div>
													</div>
												</div>
												<div class="card-body">
													<div class="card-columns">
														{{range .Nodes}}
														<div class="card">
															<div class="card-body">
																<p><strong>{{(last .Target)}}</strong></p>
																<p class="bg-secondary"><code class="text-light">{{.HTML}}</code></p>
																{{.GetPrettyFailureSummary}}
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
						</div>
					</div>
				</div>
			</div>
			{{end}}
			{{end}}
		</div>
	
		<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
		<script src="https://cdn.jsdelivr.net/npm/moment@2.24.0/moment.min.js"></script>
		<script src="https://cdn.jsdelivr.net/npm/chart.js@2.8.0/dist/Chart.min.js" integrity="sha256-Uv9BNBucvCPipKQ2NS9wYpJmi8DTOEfTA/nH2aoJALw=" crossorigin="anonymous"></script>
		<script type="text/javascript">
			var ctx = document.getElementById("resultsByOutcomePie");
			new Chart(ctx, {
				type: 'pie',
				data: {
					labels: ['Violations', 'Needs Review'],
					datasets: [{
						data: [{{.ViolationsCount}}, {{.IncompleteCount}}],
						backgroundColor: [
							'rgba(220, 53, 69, 0.2)',
							'rgba(255, 193, 7, 0.2)',
						],
						borderColor: [
							'rgba(220, 53, 69, 1)',
							'rgba(255, 193, 7, 1)',
						],
						borderWidth: 2
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

			ctx = document.getElementById("violationsBySeverityBar");
			new Chart(ctx, {
				type: 'bar',
				data: {
					labels: ['Minor', 'Moderate', 'Serious', 'Critical'],
					datasets: [{
						data: [{{.MinorCount}}, {{.ModerateCount}}, {{.SeriousCount}}, {{.CriticalCount}}],
						backgroundColor: [
							'rgba(23, 162, 184, 0.2)',
							'rgba(0, 123, 255, 0.2)',
							'rgba(255, 193, 7, 0.2)',
							'rgba(220, 53, 69, 0.2)'
						],
						borderColor: [
							'rgba(23, 162, 184, 1)',
							'rgba(0, 123, 255, 1)',
							'rgba(255, 193, 7, 1)',
							'rgba(220, 53, 69, 1)'
						],
						borderWidth: 2
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
