# AWS CloudTrail Lake data source for Grafana

The CloudTrail Lake data source plugin allows you to query and visualize AWS CloudTrail Lake logs and metrics from within Grafana.

Work in progress, with a lot of copy&paste from the [Redshift](https://github.com/grafana/redshift-datasource) and [Athena(async implementation)](https://github.com/grafana/athena-datasource) plugins.

## TODOs  & Ideas

Should have:

- understand, how to properly convert field types
    - plugin options
    - query editor
    - in TypeScript
- documentation
    - variables
    - permissions
- ship default dashboard with plugin
- clean up code
- Configuration
  - remove test options
  - allow to select default EDS
  - support Save & test

Could have:

- check other plugin types metrics, logs etc
- dashboard for filtering table
- check ad-hoc filters
- check annotation queries
- check `$__searchFilter`
- check how to [unnest JSON](https://stackoverflow.com/questions/72262711/parse-json-column-using-presto) nicely

Nice, but outside of the scope of the plugin:

- populate variable options via steampipe and permissions.cloud


## Configure the data source in Grafana

To access data source settings, hover your mouse over the **Configuration** (gear) icon, then click **Data Sources**, and then click the AWS CloudTrail Lake data source.

| Name                                | Description                                                                                                             |
| ----------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `Name`                              | The data source name. This is how you refer to the data source in panels and queries.                                   |
| `Default`                           | Default data source means that it will be pre-selected for new panels.                                                  |
| `Authentication Provider`           | Specify the provider to get credentials.                                                                                |
| `Access Key ID`                     | If `Access & secret key` is selected, specify the Access Key of the security credentials to use.                        |
| `Secret Access Key`                 | If `Access & secret key` is selected, specify the Secret Key of the security credentials to use.                        |
| `Credentials Profile Name`          | Specify the name of the profile to use (if you use `~/.aws/credentials` file), leave blank for default.                 |
| `Assume Role Arn` (optional)        | Specify the ARN of the role to assume.                                                                                  |
| `External ID` (optional)            | If you are assuming a role in another account, that has been created with an external ID, specify the external ID here. |
| `Endpoint` (optional)               | Optionally, specify a custom endpoint for the service.                                                                  |
| `Default Region`                    | Region in which the cluster is deployed.                                                                                |
| `AWS Secrets Manager`               | To authenticate with Amazon Redshift using AWS Secrets Manager.                                                         |
| `Temporary credentials`             | To authenticate with Amazon Redshift using temporary database credentials.                                              |
| `Serverless`                        | To use a Redshift Serverless workgroup.                                                                                   |
| `Cluster Identifier`                | Redshift Provisioned Cluster to use (automatically set if using AWS Secrets Manager).                                   |
| `Workgroup`                         | Redshift Serverless Workgroup to use.                                                                                   |
| `Managed Secret`                    | When using AWS Secrets Manager, select the secret containing the credentials to access the database. Note that Provisioned and Serverless stores credentials in a different format. Refer to [Storing database credentials in AWS Secrets Manager](https://docs.aws.amazon.com/redshift/latest/mgmt/data-api-access.html#data-api-secrets) for instructions. |
| `Database User`                     | User of the database. Automatically set if using AWS Secrets Manager.                                                   |
| `Database`                          | Name of the database within the cluster or workgroup.                                                                                |
| `Send events to Amazon EventBridge` | To send Data API events to Amazon EventBridge for monitoring purpose.                                                   |

## Authentication

For authentication options and configuration details, see [AWS authentication](https://grafana.com/docs/grafana/next/datasources/aws-cloudwatch/aws-authentication/) topic.

### IAM policies

Grafana needs permissions granted via IAM to be able to read Redshift metrics. You can attach these permissions to IAM roles and utilize Grafana's built-in support for assuming roles. Note that you will need to [configure the required policy](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_create.html) before adding the data source to Grafana. [You can check some predefined policies by AWS here](https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html#redshift-policy-resources.managed-policies).

Here is a minimal policy example:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AllowReadingMetricsFromRedshift",
      "Effect": "Allow",
      "Action": [
        "redshift-data:ListTables",
        "redshift-data:DescribeTable",
        "redshift-data:GetStatementResult",
        "redshift-data:DescribeStatement",
        "redshift-data:ListStatements",
        "redshift-data:ListSchemas",
        "redshift-data:ExecuteStatement",
        "redshift-data:CancelStatement",
        "redshift:GetClusterCredentials",
        "redshift:DescribeClusters",
        "redshift-serverless:ListWorkgroups",
        "redshift-serverless:GetCredentials",
        "secretsmanager:ListSecrets"
      ],
      "Resource": "*"
    },
    {
      "Sid": "AllowReadingRedshiftQuerySecrets",
      "Effect": "Allow",
      "Action": ["secretsmanager:GetSecretValue"],
      "Resource": "*",
      "Condition": {
        "Null": {
          "secretsmanager:ResourceTag/RedshiftQueryOwner": "false"
        }
      }
    }
  ]
}
```

## Query Redshift data

The provided query editor is a standard SQL query editor. Grafana includes some macros to help with writing more complex timeseries queries.

#### Macros

| Macro                        | Description                                                                                                                      | Output example                                                   |
|------------------------------|----------------------------------------------------------------------------------------------------------------------------------| ---------------------------------------------------------------- |
| `$__timeEpoch(column)`       | `$__timeEpoch` will be replaced by an expression to convert to a UNIX timestamp and rename the column to time                    | `UNIX_TIMESTAMP(dateColumn) as "time"`                           |
| `$__timeFilter(column)`      | `$__timeFilter` creates a conditional that filters the data (using `column`) based on the time range of the panel                | `time BETWEEN '2017-07-18T11:15:52Z' AND '2017-07-18T11:15:52Z'` |
| `$__timeFrom()`              | `$__timeFrom` outputs the current starting time of the range of the panel with quotes                                            | `'2017-07-18T11:15:52Z'`                                         |
| `$__timeTo()`                | `$__timeTo` outputs the current ending time of the range of the panel with quotes                                                | `'2017-07-18T11:15:52Z'`                                         |
| `$__timeGroup(column, '1m')` | `$__timeGroup` groups timestamps so that there is only 1 point for every period on the graph                                     | `floor(extract(epoch from time)/60)*60 AS "time"`                |
| `$__unixEpochFilter(column)` | `$__unixEpochFilter` be replaced by a time range filter using the specified column name with times represented as Unix timestamp | `column >= 1624406400 AND column <= 1624410000`                  |
| `$__unixEpochGroup(column)`  | `$__unixEpochGroup` is the same as $\_\_timeGroup but for times stored as Unix timestamp                                         | `floor(time/60)*60 AS "time"`                                    |
| `$__edsId()`                 | `$__edsId` outputs the selected Event Data Store's ID                                                                            | `abcdef01-2345-6789-01234-567890123456                           |

#### Table Visualization

Most queries in Redshift will be best represented by a table visualization. Any query will display data in a table. If it can be queried, then it can be put in a table.

This example returns results for a table visualization:

```sql
SELECT {column_1}, {column_2} FROM {table};
```

#### Timeseries / Graph visualizations

For timeseries / graph visualizations, there are a few requirements:

- A column with a `date` or `datetime` type must be selected
- The `date` column must be in ascending order (using `ORDER BY column ASC`)
- A numeric column must also be selected

To make a more reasonable graph, be sure to use the `$__timeFilter` and `$__timeGroup` macros.

Example timeseries query:

```sql
SELECT
  avg(execution_time) AS average_execution_time,
  $__timeGroup(start_time, 'hour'),
  query_type
FROM
  account_usage.query_history
WHERE
  $__timeFilter(start_time)
group by
  query_type,start_time
order by
  start_time,query_type ASC;
```

##### Fill value

When data frames are formatted as time series, you can choose how missing values should be filled. This in turn affects how they are rendered: with connected or disconnected values. To configure this value, change the "Fill Value" in the query editor.

#### Inspecting the query

Because Grafana supports macros that Redshift does not, the fully rendered query, which can be copy/pasted directly into Redshift, is visible in the Query Inspector. To view the full interpolated query, click the Query Inspector button, and the full query will be visible under the "Query" tab.

### Templates and variables

To add a new Redshift query variable, refer to [Add a query variable](https://grafana.com/docs/grafana/latest/variables/variable-types/add-query-variable/). Use your Redshift data source as your data source for the following available queries:

Any value queried from a Redshift table can be used as a variable. Be sure to avoid selecting too many values, as this can cause performance issues.

After creating a variable, you can use it in your Redshift queries by using [Variable syntax](https://grafana.com/docs/grafana/latest/variables/syntax/). For more information about variables, refer to [Templates and variables](https://grafana.com/docs/grafana/latest/variables/).

### Annotations

[Annotations](https://grafana.com/docs/grafana/latest/dashboards/annotations/) allow you to overlay rich event information on top of graphs. You can add annotations by clicking on panels or by adding annotation queries via the Dashboard menu / Annotations view.

**Example query to automatically add annotations:**

```sql
SELECT
  time as time,
  environment as tags,
  humidity as text
FROM
  $__table
WHERE
  $__timeFilter(time) and humidity > 95
```

The following table represents the values of the columns taken into account to render annotations:

| Name      | Description                                                                                                                       |
| --------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `time`    | The name of the date/time field. Could be a column with a native SQL date/time data type or epoch value.                          |
| `timeend` | Optional name of the end date/time field. Could be a column with a native SQL date/time data type or epoch value. (Grafana v6.6+) |
| `text`    | Event description field.                                                                                                          |
| `tags`    | Optional field name to use for event tags as a comma separated string.                                                            |

## Provision Redshift data source

You can configure the Redshift data source using configuration files with Grafana's provisioning system. For more information, refer to the [provisioning docs page](https://grafana.com/docs/grafana/latest/administration/provisioning/).

Here are some provisioning examples.

### Using AWS SDK (default)

```yaml
apiVersion: 1
datasources:
  - name: Redshift
    type: redshift
    jsonData:
      authType: default
      defaultRegion: eu-west-2
```

### Using credentials' profile name (non-default)

```yaml
apiVersion: 1

datasources:
  - name: Redshift
    type: redshift
    jsonData:
      authType: credentials
      defaultRegion: eu-west-2
      profile: secondary
```

### Using `accessKey` and `secretKey`

```yaml
apiVersion: 1

datasources:
  - name: Redshift
    type: grafana-redshift-datasource
    jsonData:
      authType: keys
      defaultRegion: eu-west-2
    secureJsonData:
      accessKey: '<your access key>'
      secretKey: '<your secret key>'
```

### Using AWS SDK Default and ARN of IAM Role to Assume

```yaml
apiVersion: 1
datasources:
  - name: Redshift
    type: grafana-redshift-datasource
    jsonData:
      authType: default
      assumeRoleArn: arn:aws:iam::123456789012:root
      defaultRegion: eu-west-2
```

## Pre-configured Redshift dashboards

Redshift data source ships with a pre-configured dashboard for some advanced monitoring parameters. This curated dashboard is based on similar dashboards in the [AWS Labs repository for Redshift](https://github.com/awslabs/amazon-redshift-monitoring). Check it out for more details.

Follow these [instructions](https://grafana.com/docs/grafana/latest/dashboards/export-import/#importing-a-dashboard) for importing a dashboard in Grafana.

Imported dashboards can be found in Configuration > Data Sources > select your Redshift data source > select the Dashboards tab to see available pre-made dashboards.

## Get the most out of the plugin

- Add [Annotations](https://grafana.com/docs/grafana/latest/dashboards/annotations/).
- Configure and use [Templates and variables](https://grafana.com/docs/grafana/latest/variables/).
- Add [Transformations](https://grafana.com/docs/grafana/latest/panels/transformations/).
- Set up alerting; refer to [Alerts overview](https://grafana.com/docs/grafana/latest/alerting/).

## Async Query Data Support

Async Query Data support enables an asynchronous query handling flow. With Async Query Data support enabled, queries will be handled over multiple requests (starting, checking its status, and fetching the results) instead of having a query be started and resolved over a single request. This is useful for queries that can potentially run for a long time and timeout.

To enable async query data support, you need to set feature toggle `redshiftAsyncQueryDataSupport` to `true`. Here are instructions to [configure feature toggles](https://grafana.com/docs/grafana/latest/setup-grafana/configure-grafana/#feature_toggles). You'll also need to ensure the IAM policy used by Grafana allows the following actions `redshift-data:ListStatements` and `redshift-data:CancelStatement`.

### Async Query Caching

To enable [query caching](https://grafana.com/docs/grafana/latest/administration/data-source-management/#query-caching) for async queries, you need to be on Grafana version 10.1 or above, and to set the feature toggles `useCachingService` and `awsAsyncQueryCaching` to `true`. You'll also need to [configure query caching](https://grafana.com/docs/grafana/latest/administration/data-source-management/#query-caching) for the specific Redshift datasource.



# original content from scaffolding

## Getting started

### Backend

1. Update [Grafana plugin SDK for Go](https://grafana.com/docs/grafana/latest/developers/plugins/backend/grafana-plugin-sdk-for-go/) dependency to the latest minor version:

   ```bash
   go get -u github.com/grafana/grafana-plugin-sdk-go
   go mod tidy
   ```

2. Build backend plugin binaries for Linux, Windows and Darwin:

   ```bash
   mage -v
   ```

3. List all available Mage targets for additional commands:

   ```bash
   mage -l
   ```
### Frontend

1. Install dependencies

   ```bash
   npm install
   ```

2. Build plugin in development mode and run in watch mode

   ```bash
   npm run dev
   ```

3. Build plugin in production mode

   ```bash
   npm run build
   ```

4. Run the tests (using Jest)

   ```bash
   # Runs the tests and watches for changes, requires git init first
   npm run test

   # Exits after running all the tests
   npm run test:ci
   ```

5. Spin up a Grafana instance and run the plugin inside it (using Docker)

   ```bash
   npm run server
   ```

6. Run the E2E tests (using Cypress)

   ```bash
   # Spins up a Grafana instance first that we tests against
   npm run server

   # Starts the tests
   npm run e2e
   ```

7. Run the linter

   ```bash
   npm run lint

   # or

   npm run lint:fix
   ```


# Distributing your plugin

When distributing a Grafana plugin either within the community or privately the plugin must be signed so the Grafana application can verify its authenticity. This can be done with the `@grafana/sign-plugin` package.

_Note: It's not necessary to sign a plugin during development. The docker development environment that is scaffolded with `@grafana/create-plugin` caters for running the plugin without a signature._

## Initial steps

Before signing a plugin please read the Grafana [plugin publishing and signing criteria](https://grafana.com/docs/grafana/latest/developers/plugins/publishing-and-signing-criteria/) documentation carefully.

`@grafana/create-plugin` has added the necessary commands and workflows to make signing and distributing a plugin via the grafana plugins catalog as straightforward as possible.

Before signing a plugin for the first time please consult the Grafana [plugin signature levels](https://grafana.com/docs/grafana/latest/developers/plugins/sign-a-plugin/#plugin-signature-levels) documentation to understand the differences between the types of signature level.

1. Create a [Grafana Cloud account](https://grafana.com/signup).
2. Make sure that the first part of the plugin ID matches the slug of your Grafana Cloud account.
   - _You can find the plugin ID in the `plugin.json` file inside your plugin directory. For example, if your account slug is `acmecorp`, you need to prefix the plugin ID with `acmecorp-`._
3. Create a Grafana Cloud API key with the `PluginPublisher` role.
4. Keep a record of this API key as it will be required for signing a plugin

## Signing a plugin

### Using Github actions release workflow

If the plugin is using the github actions supplied with `@grafana/create-plugin` signing a plugin is included out of the box. The [release workflow](./.github/workflows/release.yml) can prepare everything to make submitting your plugin to Grafana as easy as possible. Before being able to sign the plugin however a secret needs adding to the Github repository.

1. Please navigate to "settings > secrets > actions" within your repo to create secrets.
2. Click "New repository secret"
3. Name the secret "GRAFANA_API_KEY"
4. Paste your Grafana Cloud API key in the Secret field
5. Click "Add secret"

#### Push a version tag

To trigger the workflow we need to push a version tag to github. This can be achieved with the following steps:

1. Run `npm version <major|minor|patch>`
2. Run `git push origin main --follow-tags`


## Learn more

Below you can find source code for existing app plugins and other related documentation.

- [Basic data source plugin example](https://github.com/grafana/grafana-plugin-examples/tree/master/examples/datasource-basic#readme)
- [`plugin.json` documentation](https://grafana.com/developers/plugin-tools/reference-plugin-json)
- [How to sign a plugin?](https://grafana.com/docs/grafana/latest/developers/plugins/sign-a-plugin/)
