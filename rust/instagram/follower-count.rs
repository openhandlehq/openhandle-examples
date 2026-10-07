// How do I get an Instagram follower count without logging in?
// https://openhandle.dev/questions/how-do-i-get-an-instagram-follower-count-without-logging-in
// Run: cargo run --bin instagram-follower-count

use openhandle::Client;

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let response = client
        .instagram()
        .profile("northstar_forge_test")
        .get(None)
        .await?;

    // northstar_forge_test Some(48291) ...
    println!(
        "{} {:?} {:?}",
        response.data.handle, response.data.metrics.followers, response.captured_at
    );
    Ok(())
}
