import React, { useState, useEffect } from "react";
// ...existing code...
const Login = () => {
  const [showPopup, setShowPopup] = useState(false);
  const [form, setForm] = useState({ username: "", password: "" });

  useEffect(() => {
    // open popup when other code dispatches the 'openLogin' event
    const handleOpenEvent = () => setShowPopup(true);
    window.addEventListener("openLogin", handleOpenEvent);

    // also open if URL hash targets this id (optional)
    if (window.location.hash === "#login") setShowPopup(true);

    return () => window.removeEventListener("openLogin", handleOpenEvent);
  }, []);

  const handleOpen = () => setShowPopup(true);
  const handleClose = () => setShowPopup(false);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    // Handle login logic here
    console.log("Clicked submit login");
    console.log("Username:",form.username);
    console.log("Password:",form.password);
    try{
      const res=await fetch("http://localhost:8095/login", {
        method:"POST",
        headers: {
          "Content-Type":"application/json"
        },
        body:JSON.stringify({
          username:form.username,
          password:form.password
        })
      });

      if (!res.ok){
        let errText=`HTTP ${res.status}`;
        try {
          const errJson = await res.json();
          if (errJson && errJson.message) errText = errJson.message;
        } catch (_) {}
        throw new Error(errText);
      }

    const data = await res.json();
      console.log("Login success:", data);

    // example: if server returns token, store it (optional)
    if (data.token) {
      localStorage.setItem("authToken", data.token);
    }

    handleClose();

    if (data.dashboard){
      ShowDashboard(data.dashboard);
    }else{
      console.warn("No dashboard data received");
    }

  }catch(err){
    console.error("Login error:", err);
    alert(`Login failed: ${err.message}`);
  }
};

const ShowDashboard=(base64Dashboard)=>{
  if (!base64Dashboard)return;

    try {
    const binaryString = atob(base64Dashboard);
    const len = binaryString.length;
    const bytes = new Uint8Array(len);
    for (let i = 0; i < len; i++) {
      bytes[i] = binaryString.charCodeAt(i);
    }
    const blob = new Blob([bytes], { type: "application/pdf" });
    const url = URL.createObjectURL(blob);

    const newWindow = window.open(url, "_blank");
    if (!newWindow) {
      // popup blocked: trigger download/open via anchor
      const a = document.createElement("a");
      a.href = url;
      a.target = "_blank";
      a.rel = "noopener noreferrer";
      a.click();
    }

    // revoke object URL after a short delay to keep the blob valid for the opened tab
    setTimeout(() => URL.revokeObjectURL(url), 10000);
  } catch (err) {
    console.error("Failed to open dashboard PDF:", err);
    alert("Unable to open dashboard PDF.");
  }
};

  return (
    <div id="login">
      {/* keep optional programmatic open button if you want */}
      {/* <button onClick={handleOpen}>Login</button> */}
      {showPopup && (
        <div style={{
          position: "fixed",
          top: 0, left: 0, right: 0, bottom: 0,
          background: "rgba(0,0,0,0.5)",
          display: "flex", alignItems: "center", justifyContent: "center",
          zIndex: 1000
        }}>
          <div style={{
            background: "#fff",
            padding: "1rem",
            paddingBottom: "1rem",
            borderRadius: "8px",
            minWidth: "300px",
            boxShadow: "0 2px 8px rgba(0,0,0,0.2)"
          }}>
            <h2 style={{marginBottom:"1rem",
            marginTop:"0"
            }}>Login</h2>
            <form onSubmit={handleSubmit}>
              <div>
                <label>
                  Username:
                  <input
                    type="text"
                    name="username"
                    value={form.username}
                    onChange={handleChange}
                    required
                  />
                </label>
              </div>
              <div style={{ marginTop: "1rem" }}>
                <label>
                  Password:
                  <input
                    type="password"
                    name="password"
                    value={form.password}
                    onChange={handleChange}
                    required
                  />
                </label>
              </div>
              <div style={{ marginTop: "1.5rem", display: "flex", justifyContent: "space-between" }}>
                <button type="submit">Login</button>
                <button type="button" onClick={handleClose}>Cancel</button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};
export default Login;