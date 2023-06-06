# Create Virtual Environment
python3 -m venv pixelbin-sdk-venv
source pixelbin-sdk-venv/bin/activate

# Install dependencies
pip install -r requirements.txt

clear
python3 index.py

deactivate
