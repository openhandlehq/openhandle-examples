// Can I get X (Twitter) posts by date range?
// https://openhandle.dev/questions/can-i-get-x-posts-by-date-range
// Run: cargo run --bin twitter-posts-since

use chrono::{DateTime, Utc};
use openhandle::{Client, TwitterProfilePostsOptions};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    // Pagination stops once posts are older than `since`. Filter the end date yourself.
    let since: DateTime<Utc> = "2026-08-01T00:00:00Z".parse()?;
    let until: DateTime<Utc> = "2026-09-01T00:00:00Z".parse()?;

    let mut posts = client
        .twitter()
        .profile("copperfield_lab_x_test")
        .posts()
        .items(TwitterProfilePostsOptions::new().since(since));
    while let Some(post) = posts.next().await {
        let post = post?;
        if post.created_at < until {
            println!("{} {:?}", post.created_at, post.text);
        }
    }
    Ok(())
}
