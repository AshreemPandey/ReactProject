import List from './Navigation'
export default function Body(){
    return(
        <body className="body">
            <section>
                <h2>We innovate constructions<br></br>to the heights of Everest</h2>
                <ul>
                    <List value="About Us"/>
                    <List value="Contact Us"/>
                    <List value="Login"/>
                    <List value="Careers"/>
                </ul>
            </section>
        </body>
    )
}