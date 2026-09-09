import urllib.error
import urllib.request


def check_service(url="http://localhost:8080"):
    try:
        response = urllib.request.urlopen(url, timeout=5)
        print(f"Status: {response.status} OK")
        print("Response Body:\n" + response.read().decode("utf-8")[:200])
    except urllib.error.HTTPError as e:
        print(f"HTTP Error: {e.code} - {e.reason}")
    except urllib.error.URLError as e:
        print(f"Connection Failed: {e.reason}")


if __name__ == "__main__":
    check_service()