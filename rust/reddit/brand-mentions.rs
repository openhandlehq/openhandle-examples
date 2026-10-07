// How do I track brand mentions across subreddits?
// https://openhandle.dev/questions/how-do-i-track-brand-mentions-across-subreddits
// Run: cargo run --bin reddit-brand-mentions

use openhandle::{Client, RedditSearchPostsOptions};

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    // Run on a schedule. Store post IDs and skip the ones you have seen.
    let page = client
        .reddit()
        .search()
        .posts()
        .list(
            RedditSearchPostsOptions::new("synthetic")
                .sort("new")
                .t("day"),
        )
        .await?;

    for post in page.data {
        let community = post
            .community
            .and_then(|community| community.display_handle);
        println!("{:?} {:?} {:?}", community, post.metrics.score, post.title);
    }
    Ok(())
}
