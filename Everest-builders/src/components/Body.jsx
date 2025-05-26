import List from './Navigation'
import AboutUs from './AboutUs'
export default function Body(){
    return(
        <section className="Body">
            <h2>We innovate constructions<br></br>to the heights of Everest</h2>
            <ul>
                <List value="About Us" link="#AboutUs"/>
                <List value="Contact Us"/>
                <List value="Login"/>
                <List value="Careers"/>
            </ul>
        </section>
    )
}