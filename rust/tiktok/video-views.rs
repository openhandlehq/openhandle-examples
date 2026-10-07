// How do I get the view count of a TikTok video?
// https://openhandle.dev/questions/how-do-i-get-the-view-count-of-a-tiktok-video
// Run: cargo run --bin tiktok-video-views

use openhandle::{Client, Freshness, TikTokPostOptions};

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let response = client
        .tiktok()
        .post("920100000000000001")
        .get(TikTokPostOptions::new().freshness(Freshness::Live))
        .await?;

    let metrics = response.data.metrics;
    // Some(23000) Some(1200) ...
    println!(
        "{:?} {:?} {:?}",
        metrics.views, metrics.likes, response.captured_at
    );
    Ok(())
}
