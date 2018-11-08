# ASCII

This function converts text to fun ASCII art.

`POST` - `https://mca9j3f6a1.execute-api.us-east-1.amazonaws.com/dev/ascii`

Body:

```json
{ "text": "Bila shaka!" }
```

Response:

```json
{
  "text": "Bila shaka!",
  "art": "  ____    _   _                   _               _              _\n |  _ \\  (_) | |                 | |             | |            | |\n | |_) |  _  | |   __ _     ___  | |__     __ _  | | __   __ _  | |\n |  _ <  | | | |  / _` |   / __| | '_ \\   / _` | | |/ /  / _` | | |\n | |_) | | | | | | (_| |   \\__ \\ | | | | | (_| | |   <  | (_| | |_|\n |____/  |_| |_|  \\__,_|   |___/ |_| |_|  \\__,_| |_|\\_\\  \\__,_| (_)\n",
  "length": 360
}
```
