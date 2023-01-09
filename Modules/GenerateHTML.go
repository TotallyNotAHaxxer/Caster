package CastHunter

// this generates the HTML table and server data required for the server

var STDHTML = `
<!DOCTYPE html>
<html lang="en" dir="ltr">

<head>
  <meta charset="UTF-8">
  <title>Caster Dash</title>
  <link rel="stylesheet" href="style.css">
  <link href='https://unpkg.com/boxicons@2.0.7/css/boxicons.min.css' rel='stylesheet'>
  <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.8.1/css/all.css"
    integrity="sha384-50oBUHEmvpQ+1lW4y57PTFmhCaXp0ML5d60M1M7uH2+nqUivzIebhndOJK28anvf" crossorigin="anonymous">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>

<body>
  <div class="sidebar">
    <div class="logo-details">
      
      <img width="40px" height="40px" src="https://i.imgur.com/Zk8nNXn.jpg">
      <center>
        <span class="logo_name">Caster</span>
      </center>
    </div>
    <ul class="nav-links">
	<li>
	<a href="/server-index/new/indexes" class="%s">
		<i class='bx bx-ghost'></i>
		<span class="links_name">Dashboard</span>
	</a>
</li>
<li>
	<a href="/hotmaps" class="%s">
		<i class='bx bxs-hot'></i>
		<span class="links_name">Hot Settings</span>
	</a>
</li>
<li>
<a href="/devs" class="%s">
	<i class='bx bxs-devices'></i>
	<span class="links_name">Devices</span>
</a>
</li>
<li>
	<a href="/roku" class="%s">
		<i class='bx bx-cctv'></i>
		<span class="links_name">Roku Devices</span>
	</a>
</li>
<li>
	<a href="/cast" class="%s">
		<i class='bx bxl-google'></i>
		<span class="links_name">Cast Devices</span>
	</a>
</li>
<li>
<a href="/rpi" class="%s">
	<i class='bx bx-memory-card'></i>
	<span class="links_name">Rasberry PI Devices</span>
</a>
</li>
<li>
<a href="/adevs" class="%s">
  <i class='bx bxl-apple'></i>
  <span class="links_name">Apple Devices</span>
</a>
</li>
<li>
<a href="/azdevs" class="%s">
  <i class='bx bxl-amazon'></i>
  <span class="links_name">Amazon Devices</span>
</a>
</li>
<li>
<a href="/ro" class="%s">
  <i class='bx bx-wifi'></i>
  <span class="links_name">Routers</span>
</a>
</li>
<li>
	<a href="/server-index/new/indexes-errors" class="%s">
		<i class='bx bxs-error-alt'></i>
		<span class="links_name">Documentation</span>
	</a>
</li>
<li>
	<a href="/server-index/new/indexes-about" class="%s">
		<i class='bx bx-cube'></i>
		<span class="links_name">What is this tool</span>
	</a>
</li>
<li>
	<a href="https://www.github.com/ArkAngeL43" class="%s">
		<i class='bx bxl-gitlab'></i>
		<span class="links_name">Github</span>
	</a>
</li>
    </ul>
  </div>
  <section class="home-section">
    <nav>
      <div class="sidebar-button">
        <i class='bx bx-menu sidebarBtn'></i>
        <span class="dashboard">Dashboard</span>
      </div>
    </nav>
    <div class="home-content">
      <div class="overview-boxes">
        <div class="box">
          <div class="left-side">
            <div class="box-topic">IP Addresses</div>
            <hr>
            <br>
            <div class="number">%s</div>
          </div>
        </div>
        <div class="box">
          <div class="left-side">
            <div class="box-topic">Enumeration Options</div>
            <hr>
            <br>
            <div class="number">%s</div>
          </div>
        </div>
        <div class="box">
          <div class="left-side">
            <div class="box-topic">OUI's Traced</div>
            <hr>
            <br>
            <div class="number">%s</div>
          </div>
        </div>
      </div>
      <hr>
      <br>
      <div class="Info">
        <i class='bx bxs-info-square'> This page reloads every %s seconds</i>
      </div>
      <div class="table-wrapper">
        <table class="CasterTable">
          <thead>
            <tr>
              <th>IP Address</th>
              <th>MAC Address</th>
              <th>Manufacturer</th>
            </tr>
          </thead>
          <tbody>
`

var STDHTMLBOTTOM = `
<tbody>
</table>
</div>
</div>
<br>
<br>
<br>
<br>
<br>
</section>
</body>
</html>
`

var (
	Doc string
)
