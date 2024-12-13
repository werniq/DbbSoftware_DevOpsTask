from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/health', methods=['GET'])
def health():
    """
    Health check endpoint
    :return:
    """
    return jsonify(status="healthy")

if __name__ == '__main__':
    # we'd certainly won't use this in production, but for the sake of simplicity, I'd leave it here
    app.run(host='0.0.0.0')
