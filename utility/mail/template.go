package mail

var AuthHtml = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Template</title>
    <style>
        body {
            margin: 0;
            padding: 0;
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
        }

        .email-container {
            max-width: 600px;
            margin: 0 auto;
            background: #ffffff;
            border: 1px solid #e0e0e0;
            border-radius: 8px;
            overflow: hidden;
        }

        .header {
            background-color: #f4f4f4;
            color: white;
            padding: 20px;
            text-align: center;
        }

        .content {
            padding: 20px;
            color: #333333;
        }

        .footer {
            background-color: #f4f4f4;
            padding: 15px;
            text-align: center;
            font-size: 14px;
            color: #888888;
        }

        .button {
            display: inline-block;
            background-color: #8FC641;
            color: white;
            text-decoration: none;
            padding: 10px 20px;
            border-radius: 5px;
            font-size: 16px;
            margin-top: 20px;
        }

        .button:hover {
            background-color: #45a049;
        }

        .header img {
            max-width: 500px;
            margin-top: 20px;
        }
    </style>
</head>
<body>
<div class="email-container">
    <div class="header">
        <img alt="logo" src="https://res.cloudinary.com/dxjn2vrl9/image/upload/fl_preserve_transparency/v1737978983/logo_zs8ovk.png">
    </div>
    <div class="content">
        <p>Hi %s,</p>
        <p>%s</p>
        <a href=%s class="button">%s</a>
        <p>Best regards,<br>PharmaSea</p>
    </div>
    <div class="footer">
        <p>&copy; 2025 PharmaSea. All rights reserved.</p>
    </div>
</div>
</body>
</html>
`

var PharmacistHtml = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Template</title>
    <style>
        body {
            margin: 0;
            padding: 0;
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
        }

        .email-container {
            max-width: 600px;
            margin: 0 auto;
            background: #ffffff;
            border: 1px solid #e0e0e0;
            border-radius: 8px;
            overflow: hidden;
        }

        .header {
            background-color: #f4f4f4;
            color: white;
            padding: 20px;
            text-align: center;
        }

        .content {
            padding: 20px;
            color: #333333;
        }

        .footer {
            background-color: #f4f4f4;
            padding: 15px;
            text-align: center;
            font-size: 14px;
            color: #888888;
        }

        .button {
            display: inline-block;
            background-color: #8FC641;
            color: white;
            text-decoration: none;
            padding: 10px 20px;
            border-radius: 5px;
            font-size: 16px;
            margin-top: 20px;
        }

        .button:hover {
            background-color: #45a049;
        }

        .header img {
            max-width: 500px;
            margin-top: 20px;
        }
    </style>
</head>
<body>
<div class="email-container">
    <div class="header">
        <img alt="logo" src="https://res.cloudinary.com/dxjn2vrl9/image/upload/fl_preserve_transparency/v1737978983/logo_zs8ovk.png">
    </div>
    <div class="content">
        <p>Hi Dear Pharmacist,</p>
        <p>%s</p>
		<p>Email: %s</p>
		<p>Password: %s</p><br>
        <p>Best regards,<br>PharmaSea</p>
    </div>
    <div class="footer">
        <p>&copy; 2025 PharmaSea. All rights reserved.</p>
    </div>
</div>
</body>
</html>
`
