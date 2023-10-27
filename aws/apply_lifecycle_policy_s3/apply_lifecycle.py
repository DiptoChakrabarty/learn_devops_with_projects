import boto3
import sys

def apply_lifecycle_policy(bucket_name):
    # Initialize the S3 client
    s3 = boto3.client('s3')

    # Define the bucket lifecycle configuration
    lifecycle_configuration = {
        'Rules': [
            {
                'Status': 'Disabled',
                'ID': 'DeleteAfter30Days',
                'Filter': {'Prefix': ''},
                'Expiration': {'Days': 30}
            }
        ]
    }

    # Apply the lifecycle policy to the specified bucket
    s3.put_bucket_lifecycle_configuration(
        Bucket=bucket_name,
        LifecycleConfiguration=lifecycle_configuration
    )

    print(f"Lifecycle policy applied to {bucket_name}. Rules:")

    response = s3.get_bucket_lifecycle_configuration(Bucket=bucket_name)

    print(response['Rules'])

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python apply_lifecycle.py <bucket_name>")
    else:
        bucket_name = sys.argv[1]
        apply_lifecycle_policy(bucket_name)
