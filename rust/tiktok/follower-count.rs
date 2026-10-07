// How do I get a TikTok follower count for any username?
// https://openhandle.dev/questions/how-do-i-get-a-tiktok-follower-count-for-any-username
// Run: cargo run --bin tiktok-follower-count

use openhandle::Client;

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let response = client
        .tiktok()
        .profile("pixel_orchard_tt_test")
        .get(None)
        .await?;

    let profile = response.data;
    // pixel_orchard_tt_test Some(75120) Some(9814220)
    println!(
        "{} {:?} {:?}",
        profile.handle, profile.metrics.followers, profile.metrics.likes
    );
    Ok(())
}
