import Engineers from '../assets/Engineers.png'

export default function AboutUs(){
    return(
        <section className="AboutUs" id="AboutUs">
            <div className="Text">
                <p><span>Everest Builders</span> has been providing wide range of services for around <span className="experience">Fifteen</span> years.<br></br></p>

                <p>
                With quality and time as our main consideration, we convert ideas into buildings.

                We are a construction company based on Sydney. </p>
                <p>We specialize in the fields of 
                </p>
                
                <ul>
                    <li>Structures</li>
                    <li>Transportation</li>
                    <li>Irrigation</li>
                </ul>

                <p>We also offer consulting services.
                
                Our notable clients:</p>
                <ul>
                    <li>Department of Urban Planning</li>
                    <li>Macro Steels</li>
                    <li>KO mart</li>
                    <li>Brisbane Housing & Co.</li>
                </ul>
            </div>
            <img src={Engineers} alt="Two Engineers watching a map"/>

        </section>
    )
}