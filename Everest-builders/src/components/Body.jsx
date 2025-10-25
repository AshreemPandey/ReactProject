export default function Body(){
  function CareersList(e) {
    e.preventDefault();
    console.log("Clicked careers");

    fetch("http://localhost:8095/careers")
      .then((res) => {
        if (!res.ok) throw new Error("Network response was not ok");
        console.log("Fetched Careers list");
        return res.json();
      })
      .then((data) => {
          console.log("Careers:", data);

        const section = document.getElementById("careerspage");
        if (section) {
          section.scrollIntoView({ behavior: "smooth" });
        }
      })
    .catch((err) => console.error("Error:", err));
  }

  function Login(e){
    // // e.preventDefault();
    // // fetch("http://localhost:8095/login")
    // //   .then((res)=>{
    // //     if (!res.ok) throw new Error("Network response was not ok");
    // //     console.log("Fetched login form");
    // //     return res.json();
    // //   })
    //   // .then((data)=>{
    // console.log("Login clicked");
    // const section=document.getElementById("login");
    // if (section){
    //   section.scrollIntoView({behavior:"smooth"});
    // }
    //   // })
    //   // .catch((err) => console.error("Error:",err));
    // ...existing code...
    e.preventDefault();
    console.log("Login clicked");
    // const section = document.getElementById("login");
    // if (section) section.scrollIntoView({ behavior: "smooth" });

    // dispatch the event that Login.jsx listens for
    window.dispatchEvent(new Event("openLogin"));
  }
    return(
        <section className="Body">
            <h2>We innovate constructions<br></br>to the heights of Everest</h2>
            <ul>
                <a href="#AboutUs"><li>About Us</li></a>
                <a href=""><li>"Contact Us</li></a>
                <a href="#careerspage" onClick={CareersList}><li>Careers</li></a>
                <a href="#login" onClick={Login}><li>Login</li></a>
            </ul>
        </section>
    )
}