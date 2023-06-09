document.addEventListener("DOMContentLoaded", function() {
  function login() {
    var useridInput = document.querySelector("#userid-input");
    var passwordInput = document.querySelector("#password-input");
    var useridValue = useridInput.value;
    var passwordValue = passwordInput.value;


    // 예시: 간단한 유효성 검사
    if (useridValue.trim() === "" || passwordValue.trim() === "") {
      alert("아이디와 비밀번호를 입력해주세요.");
      return;
    }

    var xhr = new XMLHttpRequest();
    xhr.open('POST', 'http://localhost:8081/login', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onload = function() {
      if(xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        if(response.Success === "true") {
          // 로그인 성공
          var username = response.Username;
          showUser(username);
          loginForm.classList.add("hidden"); // 로그인 폼 숨기기
        } else {
          // 로그인 실패
          // 실패 처리 작업 필요
        }
      } else {
        console.error(xhr.statusText);
      }
    };

    xhr.onerror = function() {
      console.error(xhr.statusText);
    };


    var loginData = {
      "Userid": Number(useridValue),
      "Password": passwordValue,
      "Success": "false",
      "Username": null
    };
   
    xhr.send(JSON.stringify(loginData));
  }

  function showUser(username) {
    var usernameElement = document.getElementById("username");
    var loginButton = document.getElementById("login-button");
    var userInfo = document.getElementById("user-info");
    var logoutButton = document.getElementById("logout-button");


    usernameElement.textContent = username;
    loginButton.style.display = "none"; // signin 버튼 숨기기
    logoutButton.style.display = "block"; // signout 버튼 나타내기
    userInfo.classList.remove("hidden");
    logoutButton.addEventListener("click", logout);
  }

  // 로그인 폼 토글
  var loginButton = document.getElementById("login-button");
  var loginForm = document.getElementById("login-form");


  loginButton.addEventListener("click", function() {
    loginForm.classList.toggle("hidden");
  });


  loginForm.addEventListener("submit", function(event) {
    event.preventDefault();
    login();
  });


  // 로그아웃 버튼 클릭 시 로그아웃 처리
  var logoutButton = document.getElementById("logout-button");
  var loginForm = document.getElementById("login-form");
  var userInfo = document.getElementById("user-info");


  logoutButton.addEventListener("click", function() {
    logout();
  });


  function logout() {
    var loginButton = document.getElementById("login-button");
    var userInfo = document.getElementById("user-info");
    var logoutButton = document.getElementById("logout-button");


    userInfo.classList.add("hidden");
    logoutButton.style.display = "none"; // signout 버튼 숨기기
    loginButton.style.display = "block"; // signin 버튼 나타내기
  }


  // 과목정보 조회 버튼 클릭 시 이벤트 처리
  // 과목 정보 조회 버튼 클릭 시 이벤트 처리
  document.getElementById('subject-btn').addEventListener('click', function () {
    // AJAX 요청
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://localhost:8082/subjects', true);
    xhr.onload = function() {
      if (xhr.status === 200) {
          var subjects = JSON.parse(xhr.responseText);
          displaySubjects(subjects);
      } else {
          console.error(xhr.statusText);
      }
    };
    xhr.onerror = function() {
      console.error(xhr.statusText);
    };
    xhr.send();
  });


  // 과목 정보 표시 함수
  function displaySubjects(subjects) {
    var subjectList = document.getElementById('subject-list');
    subjectList.innerHTML = '';
  
    var table = document.createElement('table');
    table.className = 'subject-table'; // Add a class for styling purposes
  
    var tableBody = document.createElement('tbody');
  
    // Create table header row
    var headerRow = document.createElement('tr');
    headerRow.className = 'header-row'; // Add a class for styling purposes
    var headers = ['과목 ID', '과목명', '교수', '학점', '개설학과', '수강제한인원'];
    headers.forEach(function(headerText) {
      var headerCell = document.createElement('th');
      headerCell.textContent = headerText;
      headerCell.style.border = '2px solid black'; // Add border to the header cell
      headerCell.style.padding = '8px'; // Add padding to the header cell
      headerRow.appendChild(headerCell);
    });
  
    tableBody.appendChild(headerRow);
    
  
    // Create table body rows
    subjects.forEach(function(subject) {
      var row = document.createElement('tr');
  
      var rowData = [
        subject.subject_id,
        subject.subject_name,
        subject.professor,
        subject.credits,
        subject.department,
        subject.enrollment_limit
      ];
  
      rowData.forEach(function(text) {
        var cell = document.createElement('td');
        cell.textContent = text;
        cell.style.textAlign = 'center'; // Center-align the cell content
        cell.style.border = '1px solid black'; // Add border to the cell
        cell.style.padding = '6px'; // Add padding to the cell
        row.appendChild(cell);
      });
  
      tableBody.appendChild(row);
    });
  
    table.appendChild(tableBody);
    subjectList.appendChild(table);
  }
  
});
