gcloud auth activate-service-account $ACCOUNT --key-file /$KEY_FILE_NAME

appcfg.py update --application $APPLICATION --oauth2_access_token $(gcloud auth print-access-token 2> /dev/null) app/default
