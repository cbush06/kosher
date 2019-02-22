package bootstraptemplate

func GetBootstrapScaffolding() string {
	return `<html lang="en">
    <head>
        <meta charset="utf-8" />
        <title></title>
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous" />
        <link href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet" integrity="sha384-wvfXpqpZZVQGK6TAh5PVlGOfQNHSoD2xbE+QkPxCAFlNEevoEH3Sl0sibVcOQVnN" crossorigin="anonymous">
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
                            <h4 class="card-title">Features</h4>
                            <div class="ml-auto mr-auto" style="width: 400px; height: 400px;">
                                <canvas id="featuresChart" width="400" height="400"></canvas>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-6">
                    <div class="card bg-light mb-3">
                        <div class="card-body">
                            <h4 class="card-title">Scenarios</h4>
                            <div class="ml-auto mr-auto" style="width: 400px; height: 400px;">
                                <canvas id="scenariosChart" width="400" height="400"></canvas>
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

                                    <span class="col-2"><strong>Browser:</strong></span>
                                    <span class="col-4">{{.Browser}}</span>

                                    <span class="col-2"><strong>Platform:</strong></span>
                                    <span class="col-4">{{.Platform}}</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            {{.Features}}
        </div>


        <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js" integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8=" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.24.0/moment.min.js" integrity="sha256-4iQZ6BVL4qNKlQ27TExEhBN1HFPvAvAMbFavKKosSWQ=" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.7.3/Chart.min.js" integrity="sha256-oSgtFCCmHWRPQ/JmR4OoZ3Xke1Pw4v50uh6pLcu+fIc=" crossorigin="anonymous"></script>
        <script type="text/javascript">
            var ctx = document.getElementById("featuresChart");
            new Chart(ctx, {
                type: 'doughnut',
                data: {
                    labels: ['Passed', 'Failed'],
                    datasets: [{
                        data: [{{.FeaturesPassed}}, {{.FeaturesFailed}}],
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
                }
            });

            var ctx = document.getElementById("scenariosChart");
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
                }
            });
        </script>
    </body>
</html>`
}
