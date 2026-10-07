// How do I get a competitor's Instagram engagement rate?
// https://openhandle.dev/questions/how-do-i-get-a-competitors-instagram-engagement-rate
// Run: cargo run --bin instagram-engagement-rate

use openhandle::Client;

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let profile = client.instagram().profile("northstar_forge_test");
    let account = profile.get(None).await?.data;
    let page = profile.posts().list(None).await?;

    // Skip posts with hidden likes. None is not zero.
    let interactions: Vec<i64> = page
        .data
        .iter()
        .filter_map(|post| Some(post.metrics.likes? + post.metrics.comments?))
        .collect();
    let total: i64 = interactions.iter().sum();
    let followers = account.metrics.followers.unwrap_or_default();
    let rate = total as f64 / interactions.len() as f64 / followers as f64 * 100.0;

    println!("{rate:.2}% over {} posts", interactions.len());
    Ok(())
}
