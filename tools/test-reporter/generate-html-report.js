const reporter = require('cucumber-html-reporter');
const fs = require('fs');
const path = require('path');

function generateReport(options = {}) {
    const {
        jsonFile = 'reports/cucumber.json',
        output = 'reports/cucumber-html-report.html',
        projectName = 'BDD 測試專案',
        version = '1.0.0',
        environment = 'Development',
        feature = ''
    } = options;

    const workingDir = process.cwd();
    const reportsDir = path.join(workingDir, 'reports');
    
    if (!fs.existsSync(reportsDir)) {
        fs.mkdirSync(reportsDir, { recursive: true });
    }

    const reportOptions = {
        theme: 'bootstrap',
        jsonFile: jsonFile,
        output: output,
        reportSuiteAsScenarios: true,
        scenarioTimestamp: true,
        launchReport: false,
        metadata: {
            "App Version": version,
            "Test Environment": environment,
            "Browser": "N/A",
            "Platform": "Node.js",
            "Parallel": "Scenarios",
            "Executed": "Local",
            "Project": projectName
        }
    };

    if (feature) {
        reportOptions.metadata.Feature = feature;
    }

    try {
        reporter.generate(reportOptions);
        console.log(`HTML 報告已成功產生：${output}`);
        return true;
    } catch (error) {
        console.error('產生 HTML 報告時發生錯誤：', error);
        return false;
    }
}

if (require.main === module) {
    const args = process.argv.slice(2);
    const options = {};
    
    for (let i = 0; i < args.length; i += 2) {
        const key = args[i].replace('--', '');
        const value = args[i + 1];
        if (value && !value.startsWith('--')) {
            options[key] = value;
        }
    }
    
    const success = generateReport(options);
    process.exit(success ? 0 : 1);
}

module.exports = { generateReport };