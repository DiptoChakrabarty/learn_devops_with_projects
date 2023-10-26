# python script to delete s3 data

this project holds a python script which can delete S3 data which is older than 30 days

## by [@selamanse](https://github.com/selamanse)

## DESCRIPTION OF PROJECT
```
The script applies a predefined S3 bucket lifecycle policy that automatically deletes objects from the specified bucket after 30 days. You can customize the policy within the script if needed.
```

## Steps to run

### Prerequisites

Before using this script, ensure that you have the following prerequisites in place:

1. Python: Make sure you have Python installed on your system. You can download it from [python.org](https://www.python.org/downloads/).

2. Boto3: You'll need the AWS SDK for Python (Boto3). You can install it using pip:
   ```
   pip install boto3
   ```

3. AWS Credentials: Configure your AWS credentials with appropriate permissions. You can use the AWS CLI to configure your credentials:
   ```
   aws configure
   ```


### Usage

Follow these steps to use the script:

1. Clone or download the script to your local machine.

2. Open a terminal or command prompt.

3. Run the script with the following command, replacing `your-bucket-name` with the name of the target S3 bucket:

   ```
   python apply_lifecycle.py your-bucket-name
   ```

4. The script will apply the specified lifecycle policy to the bucket and display a confirmation message.


