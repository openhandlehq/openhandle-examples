// How do I search TikTok videos by keyword?
// https://openhandle.dev/questions/how-do-i-search-tiktok-videos-by-keyword
// Run: cargo run --bin tiktok-search-videos

use openhandle::{Client, TikTokSearchPostsOptions};

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let mut videos = client
        .tiktok()
        .search()
        .posts()
        .items(TikTokSearchPostsOptions::new("synthetic"));
    while let Some(video) = videos.next().await {
        let video = video?;
        let author = video.author.and_then(|author| author.handle);
        println!("{:?} {:?} {:?}", author, video.metrics.views, video.text);
    }
    Ok(())
}
