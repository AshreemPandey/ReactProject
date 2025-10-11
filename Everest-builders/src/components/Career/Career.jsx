import styles from "./Career.module.css";

export default function Careers() {
  return (
    <div className={styles.careerspage} id="careerspage">
      {/* Header Section */}
      <section className={styles.careersheader}>
        <h1>Careers</h1>

        {/* Wave Divider */}
        {/* <svg
          className={styles.wavedivider}
          viewBox="0 0 1440 320"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fill="#ffffff"
            fillOpacity="1"
            d="M0,256L60,245.3C120,235,240,213,360,197.3C480,181,600,171,720,176C840,181,960,203,1080,186.7C1200,171,1320,117,1380,90.7L1440,64L1440,320L1380,320C1320,320,1200,320,1080,320C960,320,840,320,720,320C600,320,480,320,360,320C240,320,120,320,60,320L0,320Z"
          ></path>
        </svg> */}
      </section>

      {/* Sections */}
      <Section
        title="Marketing"
        jobs={[
          "Digital Marketing Strategist",
          "SEO Specialist",
          "Brand Manager",
        ]}
      />
      <Section
        title="Development"
        jobs={[
          "Frontend Developer",
          "Backend Developer",
          "Full Stack Engineer",
        ]}
      />
      <Section
        title="Communications"
        jobs={["Content Writer", "PR Specialist", "Media Coordinator"]}
      />
    </div>
  );
}

function Section({ title, jobs }) {
  return (
    <section className={styles.careersection}>
      <h2>{title}</h2>
      <div className="divider"></div>

      <div className={styles.careerlist}>
        {jobs.map((job, idx) => (
          <div key={idx} className={styles.careercard}>
            <h3>{job}</h3>
            <p>
              We’re looking for a passionate {job.toLowerCase()} to join our
              growing team. Collaborate, innovate, and make an impact!
            </p>
            <a href="#apply" className={styles.careerbtn}>
              Read More
            </a>
          </div>
        ))}
      </div>
    </section>
  );
}
